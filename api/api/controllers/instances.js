'use strict';

const uuid = require('uuid/v1')

let instances = [
    {
        id: uuid(),
        name: "Hi"
    },
    {
        id: uuid(),
        name: "Gophers"
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
        id: uuid(),
        name: req.swagger.params.body.value.name
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
    let found_instances = instances.filter(instance => instance.id != req.swagger.params.id.value);

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

    found_instances[0].name = req.swagger.params.body.value.name;

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

        return result;
    });

    res.status(200).json(filteredInstances);
}