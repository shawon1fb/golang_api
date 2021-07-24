import { check, sleep } from "k6";
import http from "k6/http";
export let options = {
  stages: [
    { duration: "2s", target: 9000 },
    // { duration: "30s", target: 4000 },
    // { duration: "3s", target: 500 },
  ],
};

export default function () {
  let obj = {};
  obj.username = "shawon1fb";
  obj.password = "123456";
  //   let url = "http://192.168.31.121:8080/";
  let url = "https://shahanulshaheb.com/";
  let body = JSON.stringify(obj);

  let res = http.get(url);

  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
