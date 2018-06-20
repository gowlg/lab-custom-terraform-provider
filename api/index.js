'use strict';

const SwaggerExpress = require('swagger-express-mw');
const swaggerUiAssetPath = require("swagger-ui-dist").getAbsoluteFSPath()
const express = require('express');

const port = process.env.PORT || 3000;
const config = {
    appRoot: __dirname
};

const app = express()

SwaggerExpress.create(config, function(err, swaggerExpress) {
    if (err) { throw err; }

    app.set('view engine', 'ejs');

    app.get('/docs', function(req, res) {
        res.render('swaggerUI');
    });

    app.use('/docs/assets', express.static(swaggerUiAssetPath));

    swaggerExpress.register(app);
});

module.exports = app.listen(port);