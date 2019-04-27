import axios from 'axios'
import Vue from 'vue'
import VueCookie from 'vue-cookie'

Vue.use(VueCookie)

export const HTTPAUTH = axios.create({
  baseURL: process.env.AUTH_HOST,
  headers: {
    'Content-Type': 'application/json',
    'Authorization': getAuthToken()
  }
})

export const HTTP = axios.create({
  baseURL: process.env.API_HOST,
  headers: {
    'Content-Type': 'application/json',
    'Authorization': getAuthToken()
  }
})

export function getToken() {
  return getAuthData()
}

export function clearToken() {
    Vue.cookie.delete('rest-doc')
}

function getAuthToken() {
    var auth = getAuthData()
    if (auth) {
        return auth.token_type + " " + auth.access_token;
    }
    return "";
}

function getAuthData() {
    const auth = Vue.cookie.get('rest-doc')
    const authObj = JSON.parse(auth)
    var ts = Math.round((new Date()).getTime() / 1000);
    if (authObj && (ts < authObj.expires_in)) {
      return authObj
    }
    return null
}
