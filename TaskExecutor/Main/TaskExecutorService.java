import java.util.LinkedList;
import java.util.Map;
import java.util.Queue;
import java.util.UUID;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

public class TaskExecutorService implements TaskExecutor {
        private final ExecutorService executorService;
        private final Map<UUID, Object> taskGroupLocks = new ConcurrentHashMap();
        private final Queue<Runnable> taskQueue = new LinkedList<>();
        private final Object queueLock = new Object();

        public TaskExecutorService(int maxConcurrency) {
            this.executorService = Executors.newFixedThreadPool(maxConcurrency);
            startQueueProcessor();
        }

        @Override
        public <T> Future<T> submitTask(Task<T> task) {
            CompletableFuture<T> future = new CompletableFuture<>();

            synchronized (queueLock) {
                taskQueue.add(() -> executeTask(task, future));
                queueLock.notify();
            }

            return future;
        }

        private <T> void executeTask(Task<T> task, CompletableFuture<T> future) {
            Object lock = taskGroupLocks.computeIfAbsent(task.taskGroup().groupUUID(), k -> new Object());

            synchronized (lock) {
                try {
                    T result = task.taskAction().call();
                    future.complete(result);
                } catch (Exception e) {
                    future.completeExceptionally(e);
                }
            }
        }

        private void startQueueProcessor() {
            Thread processorThread = new Thread(() -> {
                while (true) {
                    Runnable task;

                    synchronized (queueLock) {
                        while (taskQueue.isEmpty()) {
                            try {
                                queueLock.wait();
                            } catch (InterruptedException e) {
                                Thread.currentThread().interrupt();
                                return;
                            }
                        }

                        task = taskQueue.poll();
                    }

                    executorService.submit(task);
                }
            });

            processorThread.setDaemon(true);
            processorThread.start();
        }

        public void shutdown() {
            executorService.shutdown();
        }

        public static void main(String[] args) throws ExecutionException, InterruptedException {
            TaskExecutorService taskExecutor = new TaskExecutorService(4);

            TaskGroup group1 = new TaskGroup(UUID.randomUUID());
            TaskGroup group2 = new TaskGroup(UUID.randomUUID());

            Task<Integer> task1 = new Task<>(UUID.randomUUID(), group1, TaskType.READ, () -> {
                Thread.sleep(1000);
                return 1;
            });

            Task<Integer> task2 = new Task<>(UUID.randomUUID(), group1, TaskType.WRITE, () -> {
                Thread.sleep(500);
                return 2;
            });

            Task<Integer> task3 = new Task<>(UUID.randomUUID(), group2, TaskType.READ, () -> {
                Thread.sleep(300);
                return 3;
            });

            Future<Integer> future1 = taskExecutor.submitTask(task1);
            Future<Integer> future2 = taskExecutor.submitTask(task2);
            Future<Integer> future3 = taskExecutor.submitTask(task3);

            System.out.println("Task 1 result: " + future1.get());
            System.out.println("Task 2 result: " + future2.get());
            System.out.println("Task 3 result: " + future3.get());

            taskExecutor.shutdown();
        }
    }
