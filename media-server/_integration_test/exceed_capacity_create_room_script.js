import http from "k6/http";
import { check } from "k6";
import { Rate } from "k6/metrics";

const failureRate = new Rate("check_failure_rate");

export default function() {
    var rooms = 10000;
    for(let i = 0; i < rooms; i++) {
        let createRoomResponse = http.post("http://host.docker.internal:3000/api/v1/room/");
        let checkCreateRoomRes = check(createRoomResponse, {
            "created status is 201": (r) => r.status === 201,
            "created roomId": r => r.json()["roomId"] !== undefined,
            "created serverId": r => r.json()["serverId"] !== undefined,
        });
        failureRate.add(!checkCreateRoomRes);
    }
};