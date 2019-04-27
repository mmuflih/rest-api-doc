<template lang="html">
    <div class="login-box">
        <div class="login-logo">
            <router-link :to="{path: '/'}"><b>Vue</b>Admin LTE</router-link>
        </div>
        <div class="login-box-body">
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
                <div class="row">
                    <div class="col-xs-4">
                        <button type="submit" class="btn btn-primary btn-block btn-flat">Sign In</button>
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
                password: ''
            }
        }
    },
    methods: {
        login() {
            HTTP.post("/api/v1/user/login", this.data).then(response => {
                this.$cookie.set('rrauth', JSON.stringify(response.data.data))
                location.href = "/"
            }).catch(error => {
                this.$toasted.error(error.response.data.user_message).goAway(3000)
            });
        }
    }
}
</script>

<style lang="css">
</style>
