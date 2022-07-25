var viewer = new Cesium.Viewer("cesiumContainer");

// add geoserver wms layer
var provider = new Cesium.WebMapServiceImageryProvider({
    url: 'http://localhost:8080/geoserver/topp/wms',
    layers: 'topp:states',
    parameters: {
        service : 'WMS',
        format: 'image/png',
        transparent: true,
    }
    }
);

viewer.imageryLayers.addImageryProvider(provider);