'use strict';

let idGenerator = 0
function getId() {
    return idGenerator++;
}

let instances = [
    {
        id: getId(),
        name: "Hi",
        memory: 300,
        type: "t2.nano",
        private: false,
        tag: "Srizzling rulez"
    },
    {
        id: getId(),
        name: "Gophers",
        memory: 500,
        type: "t2.large",
        private: true
    }
]

module.exports = {
    create_instance: create_instance,
    read_instance: read_instance,
    update_instance: update_instance,
    delete_instance: delete_instance,
    list_instances: list_instances,
};

function create_instance(req, res) {
    const newInstance = {
        id: getId(),
        name: req.swagger.params.body.value.name,
        memory: req.swagger.params.body.value.memory,
        type: req.swagger.params.body.value.type,
        private: req.swagger.params.body.value.private,
        tag: req.swagger.params.body.value.tag,
    };

    instances.push(newInstance);

    res.status(201).json(newInstance);
}

function read_instance(req, res) {
    let found_instances = instances.filter(instance => instance.id == req.swagger.params.id.value);

    if(found_instances.length > 1) {
        res.status(500).json({
            message: "Error: Expected 0 "
        });

        return;
    }
    
    if(found_instances.length < 1) {
        res.status(404).json({
            message: "Unable to find an instance with that ID"
        });

        return;
    }

    res.status(200).json(found_instances[0]);
}

function update_instance(req, res) {
    let found_instances = instances.filter(instance => instance.id == req.swagger.params.id.value);

    if(found_instances.length > 1) {
        res.status(500).json({
            message: `Error: Expected 1 but got ${found_instances.length}`
        });

        return;
    }
    
    if(found_instances.length < 1) {
        res.status(404).json({
            message: "Unable to find an instance with that ID"
        });

        return;
    }

    found_instances[0].name = req.swagger.params.body.value.name;
    found_instances[0].memory = req.swagger.params.body.value.memory;
    found_instances[0].type = req.swagger.params.body.value.type;
    found_instances[0].private = req.swagger.params.body.value.private;
    found_instances[0].tag = req.swagger.params.body.value.tag;

    res.status(200).json(found_instances[0]);
}

function delete_instance(req, res) {
    instances = instances.filter(instance => instance.id != req.swagger.params.id.value);

    res.status(200).json({
        message: "Instance successfully deleted"
    });
}

function list_instances(req, res) {
    let filteredInstances = instances.filter(instance => {
        let result = true;
        
        if(req.swagger.params.id.value) {
            result = result && req.swagger.params.id.value == instance.id;
        }

        if(req.swagger.params.name.value){
            result = result && req.swagger.params.name.value == instance.name;
        }

        if(req.swagger.params.memory.value){
            result = result && req.swagger.params.memory.value == instance.memory;
        }

        if(req.swagger.params.type.value){
            result = result && req.swagger.params.type.value == instance.type;
        }

        if(req.swagger.params.tag.value){
            result = result && req.swagger.params.tag.value == instance.tag;
        }

        if(req.swagger.params.private.value != null){
            result = result && req.swagger.params.private.value == instance.private;
        }

        return result;
    });

    res.status(200).json(filteredInstances);
}