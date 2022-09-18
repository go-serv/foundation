import {grpc} from "@improbable-eng/grpc-web";

// Import code-generated data structures.
import {NetParcel} from "../autogen/net/net_pb_service.js";
import * as Ping from "../autogen/net/ping_pb.js";

const pingReq = new Ping.Ping.Request();
pingReq.setPayload(3)

grpc.invoke(NetParcel.Ping, {
    request: pingReq,
    host: "https://0.0.0.0:3033",
    onMessage: (message: Ping.Ping.Response) => {
        console.log("pong: " + message.getPayload());
    },
    onEnd: (code: grpc.Code, msg: string | undefined, trailers: grpc.Metadata) => {
        if (code == grpc.Code.OK) {
            console.log("all ok")
        } else {
            console.log("hit an error", code, msg, trailers);
        }
    }
});