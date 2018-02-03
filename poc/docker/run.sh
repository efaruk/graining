docker_host="localhost"
name="autoc"
tag="autoc"

docker -H $docker_host run -d -e "BASIC_ENV_VARIABLE_3=Running as Production" --cpu-quota=200000 -m 5GB --name $name $tag
