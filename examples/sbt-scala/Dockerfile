# Copyright 2016 Cisco Systems, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
# http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM library/java:8-jre

COPY target/scala-2.11/hello-akka-assembly-1.0.jar /opt/app.jar

COPY ./init.d /opt/init.d
WORKDIR /opt/service

EXPOSE 8080

ENTRYPOINT ["/opt/init.d/server-entrypoint.sh"]
