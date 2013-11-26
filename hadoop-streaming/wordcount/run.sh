#!/bin/sh
 
###############################################
 
jobname="word count"
hadoop_in="/user/test/input"
hadoop_out="/user/test/output"
 
###############################################
 
currentPath=`pwd`;
mapper=$currentPath"/wordcount_map.rb"
reducer=$currentPath"/wordcount_reduce.rb"

###############################################
 
options="-D mapred.reduce.tasks=30"
options=$options" -D mapred.job.priority=NORMAL"
#options=$options" -file "$currentPath"/dic.json"
 
###############################################
 
echo "remove hdfs file: \""$hadoop_out"\" ?[y/n]"
read ANS
 
if [ $ANS = 'y' -o $ANS = 'yes' ]; then
	hadoop fs -rmr ${hadoop_out}
	
	hadoop jar /usr/lib/hadoop/contrib/streaming/hadoop-streaming-0.20.2-cdh3u6.jar \
		-mapper  ${mapper} \
		-reducer ${reducer} \
		-input   ${hadoop_in} \
		-output  ${hadoop_out} \
		-file    ${mapper} \
		-file    ${reducer}
fi
