# Getting Started with Cloud Spanner and the Google Cloud Client libraries

<a href="https://console.cloud.google.com/cloudshell/open?git_repo=https://github.com/googleapis/java-spanner&page=editor&open_in_editor=samples/README.md">
<img alt="Open in Cloud Shell" src ="http://gstatic.com/cloudssh/images/open-btn.png"></a>

[Cloud Spanner][Spanner] is a horizontally-scalable database-as-a-service
with transactions and SQL support.
These sample Java applications demonstrate how to access the Spanner API using
the [Google Cloud Client Library for Java][java-spanner].

[Spanner]: https://cloud.google.com/spanner/
[java-spanner]: https://github.com/googleapis/java-spanner

## Quickstart

Install [Maven](http://maven.apache.org/).

Build your project from the root directory (`java-spanner`):

    mvn clean package -DskipTests -DskipUTs -Penable-samples

Every subsequent command here should be run from a subdirectory (`cd samples/snippets`).

You can run a given `ClassName` via:

    mvn exec:java -Dexec.mainClass=com.example.spanner.ClassName \
        -DpropertyName=propertyValue \
        -Dexec.args="any arguments to the app"

### Running a simple query (using the quickstart sample)

    mvn exec:java -Dexec.mainClass=com.example.spanner.QuickstartSample -Dexec.args="my-instance my-database"

## Tutorial

### Running the tutorial
    mvn exec:java -Dexec.mainClass=com.example.spanner.SpannerSample -Dexec.args="<command> my-instance my-database"

## Tracing sample
`TracingSample.java` demonstrates how to export traces generated by client library to StackDriver and to /tracez page.

### Running the tracing sample
    mvn exec:java -Dexec.mainClass=com.example.spanner.TracingSample -Dexec.args="my-instance my-database"

## Test
    mvn verify -Dspanner.test.instance=<instance id> -Dspanner.sample.database=<new database id>  -Dspanner.quickstart.database=<existing database id>