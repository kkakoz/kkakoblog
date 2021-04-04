import axios from "axios";
import QS from "qs";

axios.defaults.baseURL = "http://127.0.0.1:10003";
axios.defaults.timeout = 10000;
axios.defaults.headers.post["Content-Type"] =
  "application/x-www-form-urlencoded;charset=UTF-8";

var instance = axios.create({
  baseURL: "http://127.0.0.1:10003",
  timeout: 10000
});

export function get(url, params) {
  return new Promise((resolve, reject) => {
    instance
      .get(url, {
        params: params
      })
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err.data);
      });
  });
}

export function post(url, params) {
  return instance.post(url, QS.stringify(params));
}

export const BlogLogin = p => post("blog/user/login", p);

export const GithubConf = p => get("blog/oauth/github_conf", p);

export const OAuthLogin = p => post("blog/oauth/github_auth", p);
