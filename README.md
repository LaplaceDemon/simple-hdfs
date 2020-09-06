# simple-hdfs
Simple hdfs command line tool.


## Build binary

On Windows:
```
build.bat
```

On Linux/Mac OS:
```
chmod +x build.sh
build.sh
```


## Commands

Create a directory

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username mkdir /test

simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username mkdirall /test/hdfs
```


Create a file

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username touch /test/myfile.txt
```

List of files and directories

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username ls /test

simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username ll /test
```


Delete a file

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username rm /test/myfile.txt
```

Delete a directory

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username rmdir /test/hdfs
```


Delete a non-empty directory

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username rmdirall /test
```


Rename a file

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username rename /test/myfile.txt /test/myfile-00.txt
```


Read the text data of the front part of the file. By default, 32k length data will be read. You can also specify the read length

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username less /test/myfile.txt

simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username less /test/myfile.txt 512
```


Read the data of any part of the file. Need to specify the offset and length.

```bash
simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username readAt /test/myfile.txt 0 1024

simple-hdfs -nn0=name_node_0_host:port -nn1=name_node_1_host:port -user=username readAt /test/myfile.txt 512 1024
```