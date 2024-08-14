# PRAVAH

**This is a light weight workflow execution engine written in GO**

## **Project status**

- - - -

* Still in development, not recommended to use in production
* More features to be added, in pipeline.

## **Documentation**

- - - -

### Documentations

* <https://fssplatform.atlassian.net/wiki/spaces/PEDOC/pages/1442519/Process+Engine>
* <https://fssplatform.atlassian.net/wiki/spaces/PEDOC/pages/84312069/Criteria+of+implementing+a+BPM+methodology>
* <https://fssplatform.atlassian.net/wiki/spaces/PEDOC/pages/235143269/Audit+Trail>


## **Requirements**

- - - -

Go 1.20+

## **BPMN Modelling**

- - - -

BPMN Models for Testing:
<https://bitbucket.org/fsstech/pravah/src/04c7445ede927541c57d4f560f2034c7619e293c/BPMN/?at=feature%2FdeployFlowExceptionHandling>

All these examples are build with [Camunda Modeler Community Edition](https://camunda.com/de/download/modeler/). We have made the diagrams on Camunda Platform 8. I would like to send a big "thank you", to Camunda for providing such tool.

## **Implementation notes**

- - - -

IDs (process definition, process instance, job, events, etc.)

This engine does use an implementation of Twitter's Snowflake algorithm which combines some advantages, like it's time based and can be sorted, and it's collision free to a very large extend. So you can rely on larger IDs were generated later in time, and they will not collide with IDs, generated on e.g. other nodes of your application in a multi-node installation.
The IDs are structured like this ...

```
+-----------------------------------------------------------+
| 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+-----------------------------------------------------------+
```

The NodeID is generated out of a hash-function which reads all environment variables. As a result, this approach allows 4096 unique IDs per node and per millisecond.

# test
