FROM ubuntu:focal
RUN apt update
RUN apt upgrade -y
# install and run nginx
RUN env DEBIAN_FRONTEND=noninteractive apt install nginx -y
ADD docker_confs/nginx_config /
ADD dev_tools/start_script.sh /usr/local/bin/start_script.sh
RUN chmod +x /usr/local/bin/start_script.sh
CMD ["start_script.sh"]