package com.example.mysql.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.mysql.ResourceNotFoundException;
import com.example.mysql.model.Comment;
import com.example.mysql.repository.CommentRepository;

@RestController
@RequestMapping("/api")
public class CommentController {

	@Autowired
	CommentRepository commentRepository;

	@GetMapping("/comments")
	public List<Comment> getAllComments() {
		return commentRepository.findAll();
	}

	@PostMapping("/comments")
	public Comment createComment(@RequestBody Comment comment) {
		return commentRepository.save(comment);
	}

	@GetMapping("/comments/{id}")
	public Comment getCommentById(@PathVariable(value = "id") Long commentId) {
		return commentRepository.findById(commentId)
				.orElseThrow(() -> new ResourceNotFoundException("Comment", "id", commentId));
	}

	@PutMapping("/comments/{id}")
	public Comment updateComment(@PathVariable(value = "id") Long commentId, @RequestBody Comment commentDetails) {

		Comment comment = commentRepository.findById(commentId)
				.orElseThrow(() -> new ResourceNotFoundException("Comment", "id", commentId));

		comment.setPostName(commentDetails.getPostName());
		comment.setComment(commentDetails.getComment());

		Comment updatedComment = commentRepository.save(comment);
		return updatedComment;
	}

	@DeleteMapping("/comments/{id}")
	public ResponseEntity<?> deleteComment(@PathVariable(value = "id") Long commentId) {
		Comment comment = commentRepository.findById(commentId)
				.orElseThrow(() -> new ResourceNotFoundException("Comment", "id", commentId));

		commentRepository.delete(comment);

		return ResponseEntity.ok().build();
	}
}
