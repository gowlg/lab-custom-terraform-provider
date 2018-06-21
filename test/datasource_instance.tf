data "uws_instance" "my_instance" {
  instance_id = 1
}

output "datasource_uws_instance_instance_id" {
  value = "${data.uws_instance.my_instance.instance_id}"
}

output "datasource_uws_instance_name" {
  value = "${data.uws_instance.my_instance.name}"
}

output "datasource_uws_instance_memory" {
  value = "${data.uws_instance.my_instance.memory}"
}

output "datasource_uws_instance_type" {
  value = "${data.uws_instance.my_instance.type}"
}

output "datasource_uws_instance_tag" {
  value = "${data.uws_instance.my_instance.tag}"
}

output "datasource_uws_instance_private" {
  value = "${data.uws_instance.my_instance.private}"
}
