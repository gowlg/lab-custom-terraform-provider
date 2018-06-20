data "uws_instance" "my_instance" {
  instance_id = 0
}

output "instance_id" {
  value = "${data.uws_instance.my_instance.instance_id}"
}

output "name" {
  value = "${data.uws_instance.my_instance.name}"
}

output "memory" {
  value = "${data.uws_instance.my_instance.memory}"
}

output "type" {
  value = "${data.uws_instance.my_instance.type}"
}

output "tag" {
  value = "${data.uws_instance.my_instance.tag}"
}

output "private" {
  value = "${data.uws_instance.my_instance.private}"
}
