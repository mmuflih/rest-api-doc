<template lang="html">
    <div class="login-box">
        <div class="login-logo">
            <router-link :to="{path: '/'}"><b>Rest</b>ful DOC</router-link>
        </div>
        <div class="login-box-body" v-if="!registerNew">
            <p class="login-box-msg">Sign in</p>
            <form v-on:submit.prevent="login">
                <div class="form-group has-feedback">
                    <input type="email" class="form-control" placeholder="Email" v-model="data.email">
                    <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback">
                    <input type="password" class="form-control" placeholder="Password" v-model="data.password">
                    <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                </div>
                <div class="col-md-12">
                    <div class="col-md-6">
                        <button type="submit" class="btn btn-primary btn-block btn-flat">Sign In</button>
                    </div>
                    <div class="col-md-6">
                        <button @click="register" class="btn btn-primary btn-block btn-flat">Sign Up</button>
                    </div>
                </div>
            </form>
        </div>
        <div class="login-box-body" v-if="registerNew">
            <p class="login-box-msg">Sign in</p>
            <form v-on:submit.prevent="registerSubmit">
                <div class="form-group has-feedback">
                    <input type="email" class="form-control" placeholder="Email" v-model="data.email">
                    <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback">
                    <input type="text" class="form-control" placeholder="Name" v-model="data.name">
                    <span class="glyphicon glyphicon-user form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback">
                    <input type="text" class="form-control" placeholder="Phone" v-model="data.phone">
                    <span class="glyphicon glyphicon-phone form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback">
                    <input type="password" class="form-control" placeholder="Password" v-model="data.password">
                    <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                </div>
                <div class="col-md-12">
                    <div class="col-md-6">
                        <button class="btn btn-primary btn-block btn-flat" @click="doCancel">Batal</button>
                    </div>
                    <div class="col-md-6">
                        <button type="submit" class="btn btn-primary btn-block btn-flat">Register</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
import {HTTP} from '@/lib/httplib'

export default {
    data() {
        return {
            data: {
                email: '',
                password: '',
            },
            registerNew: false
        }
    },
    methods: {
        login() {
            HTTP.post("/api/v1/user/login", this.data).then(response => {
                this.$cookie.set('rest-doc', JSON.stringify(response.data.data))
                location.href = "/"
            }).catch(error => {
                this.$toasted.error(error.response.data.user_message).goAway(3000)
            });
        },
        register() {
            this.registerNew = true
        },
        doCancel() {
            this.registerNew = false
        },
        registerSubmit() {
            HTTP.post("/api/v1/user/login", this.data).then(response => {
                this.$cookie.set('rest-doc', JSON.stringify(response.data.data))
                location.href = "/"
            }).catch(error => {
                this.$toasted.error(error.response.data.user_message).goAway(3000)
            });
        }
    }
}
</script>

<style lang="css">
    .login-box-body {
        border: 1px solid #000;
    }
</style>
