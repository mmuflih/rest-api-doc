import axios from 'axios'
import Vue from 'vue'
import VueCookie from 'vue-cookie'

export function http(hearders) {
    return axios.create({
        headers: headers
    })
}

