import { check, sleep } from "k6";
import http from "k6/http";
export let options = {
  stages: [
    { duration: "20s", target: 100 },
    { duration: "30s", target: 1000 },
    { duration: "3s", target: 10000 },
  ],
};
export default function () {
  let obj = {};
  obj.username = "shawon1fb";
  obj.password = "123456";
  let url = "http://192.168.31.121:8080/users/login";
  let body = JSON.stringify(obj);

  let res = http.post(url, body);

  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
