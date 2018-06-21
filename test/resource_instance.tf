resource "uws_instance" "new_instance" {
  name    = "Test Instance"
  memory  = 200
  type    = "t2.unicorn"
  tag     = "Alex Sucks2"
  private = false
}

output "resource_uws_instance_instance_id" {
  value = "${uws_instance.new_instance.instance_id}"
}

output "resource_uws_instance_name" {
  value = "${uws_instance.new_instance.name}"
}

output "resource_uws_instance_memory" {
  value = "${uws_instance.new_instance.memory}"
}

output "resource_uws_instance_type" {
  value = "${uws_instance.new_instance.type}"
}

output "tag" {
  value = "${uws_instance.new_instance.tag}"
}

output "private" {
  value = "${uws_instance.new_instance.private}"
}
