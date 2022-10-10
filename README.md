# cloud-function-poc
This is a sample project to try out a POC for cloud functions written in go deployed to GCP

**Prerequisites**

* Java & Maven
* Go
* Make

This project has 3 folders

**assembly** 

    assembly
    | -- function_1.xml 
    | -- function_2.xml
    | -- ...
    | -- ...

This folder includes all the xml files for assembly and packaging of zip of the individual cloud function


**cloud-function**

    cloud-functions
    | -- function_1
            | -- function_1.go
            | -- ...
            | -- ...
    | -- function_2
            | -- function_2.go
            | -- ...
            | -- ...

This folder contains actual source code per cloud function

**common**

    common
    | -- helper
            | -- sum.go
            | -- ...
            | -- ...
    | -- logging
            | -- logger.go
            | -- ...
            | -- ...

This folder contains all the common code, helper code which is required by almost all the cloud functions in the project.
This can contain multiple folders under the common folder and common go files can go under the common folder directly.


**Assembling the zip for a cloud function**

For assembling the cloud function into an individual zip, you can refer the file

`function_assembly_template.xml`

    <fileSet>
            <directory>${project.basedir}/common</directory>
            <includes>
                <include>**/**.go</include>
            </includes>
            <excludes>
                <exclude>**/*_test.go</exclude>
            </excludes>
    </fileSet>
    <fileSet>
        <directory>${project.basedir}/cloud-functions/FUNCTION_NAME</directory>
        <outputDirectory>/</outputDirectory>
        <includes>
            <include>**/**.go</include>
        </includes>
        <excludes>
            <exclude>**/**_test.go</exclude>
        </excludes>
    </fileSet>


For packaging the cloud function you can specify all the files which you want to include and exclude in the xml. 
After the things are finalised put the final xml into the assembly folder and you shall have your zip build after you build the project


**How to build this project**

`make build`
This command will

1. clean
2. download dependency
3. format
4. compile
5. test
6. build zips

The final zips will be present in the target folder