import {grpc} from "@improbable-eng/grpc-web";
// Import code-generated data structures.
import {NetParcel} from "../autogen/net/net_pb_service";
import * as Ping from "../autogen/net/ping_pb";
import {describe, expect, test} from "@jest/globals";

const PingPayload = 7

// @ts-ignore
function invoke(resolve, reject) {
    const pingReq = new Ping.Ping.Request();
    pingReq.setPayload(PingPayload)
    grpc.invoke(NetParcel.Ping, {
        request: pingReq,
        host: "http://go-serv.io:3033",
        onMessage: (message: Ping.Ping.Response) => {
            resolve(message)
        },
        onEnd: (code: grpc.Code, msg: string | undefined, trailers: grpc.Metadata) => {
            if (code != grpc.Code.OK) {
                reject(code)
            }
        }
    });
}

describe('NetParcel service calls', () => {
    test('ping', async () => {
        return new Promise<any>((resolve, reject) => {
            invoke(resolve, reject)
        }).then((e: Ping.Ping.Response) => {
            expect(e.getPayload()).toBe(PingPayload)
        }).catch((e) => {
            throw new Error(e)
        })
    });
});