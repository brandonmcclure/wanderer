import { ValhallaRoute, type ValhallaResponse } from "$lib/models/valhalla";
import { ClientResponseError } from "pocketbase";



export const route: ValhallaRoute = new ValhallaRoute();


export function clearRoute() {
    route.geometry.coordinates = [];
}

export async function calculateRouteBetween(startLat: number, startLon: number, endLat: number, endLon: number) {
    const requestBody = {
        "directions_type": "none",
        "locations": [{ "lat": startLat, "lon": startLon }, { "lat": endLat, "lon": endLon }],
        "costing": "pedestrian", "costing_options": { "pedestrian": { "max_hiking_difficulty": 6, "use_ferry": 0 } }
    }
    const r = await fetch("/api/v1/valhalla", { method: "POST", body: JSON.stringify(requestBody) })
    if (!r.ok) {
        throw new ClientResponseError(await r.json())
    }
    const response: ValhallaResponse = await r.json();

    appendToRoute(response);
}

function appendToRoute(valhallaResponse: any) {
    const routeGeometry = decodeShape(valhallaResponse.trip.legs[0].shape);

    for (const point of routeGeometry) {
        route.geometry.coordinates.push([point[1], point[0]])
    }
}


function decodeShape(shape: string, precision: number = 6) {
    var index = 0,
        lat = 0,
        lng = 0,
        coordinates = [],
        shift = 0,
        result = 0,
        byte = null,
        latitude_change,
        longitude_change,
        factor = Math.pow(10, precision);

    while (index < shape.length) {
        byte = null;
        shift = 0;
        result = 0;

        do {
            byte = shape.charCodeAt(index++) - 63;
            result |= (byte & 0x1f) << shift;
            shift += 5;
        } while (byte >= 0x20);

        latitude_change = ((result & 1) ? ~(result >> 1) : (result >> 1));

        shift = result = 0;

        do {
            byte = shape.charCodeAt(index++) - 63;
            result |= (byte & 0x1f) << shift;
            shift += 5;
        } while (byte >= 0x20);

        longitude_change = ((result & 1) ? ~(result >> 1) : (result >> 1));

        lat += latitude_change;
        lng += longitude_change;

        coordinates.push([lat / factor, lng / factor]);
    }

    return coordinates;
};