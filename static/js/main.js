$(function() {
    var url = "data/swagger.json";
    //var url = "http://petstore.swagger.io/v2/swagger.json";

    hljs.configure({
        highlightSizeThreshold: 5000
    });

    // Pre load translate...
    if (window.SwaggerTranslator) {
        window.SwaggerTranslator.translate();
    }
    window.swaggerUi = new SwaggerUi({
        url: url,
        dom_id: "swagger-ui-container",
        onComplete: function(swaggerApi, swaggerUi) {
            // set page title
            document.title = $('div.info_title').text();
        },
        onFailure: function(data) {
            log("Unable to Load SwaggerUI");
        },
        docExpansion: "list",
        jsonEditor: false,
        defaultModelRendering: 'schema',
        showRequestHeaders: false,
        operationsSorter: 'method',
        validatorUrl: null,
        supportedSubmitMethods: ['']
    });

    window.swaggerUi.load();

    function log() {
        if ('console' in window) {
            console.log.apply(console, arguments);
        }
    }
});