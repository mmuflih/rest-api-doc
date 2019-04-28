<template>
    <div class="col-md-12">
        <div class="box">
            <div class="overlay" v-if="loading">
                <i class="fa fa-gear fa-spin"></i>
            </div>
            <div class="box-header">
                <div class="box-title"><h4>Title</h4></div>
                <div class="pull-right">
                    <select class="form-control" v-model="endpoint.env">
                        <option value="https://sandbox-api.selenago.com">Sandbox</option>
                        <option value="http://localhost:8055">Local</option>
                        <option value="https://api.selenago.com">Production</option>
                    </select>
                </div>
            </div>
            <div class="box-body">
                <div class="form-group">
                    <label for="endpoint">Endpoint</label>
                    <div class="form-inline">
                        <div class="col-md-12">
                            <div class="row">
                                <div class="col-md-1">
                                    <div class="row">
                                        <select class="form-control">
                                            <option value="POST">POST</option>
                                            <option value="GET">GET</option>
                                            <option value="PUT">PUT</option>
                                            <option value="DELETE">DELETE</option>
                                        </select>
                                    </div>
                                </div>
                                <div class="col-md-8">
                                    <div class="row">
                                        <input class="form-control" id="endpoint" style="width: 100%;" v-model="endpoint.url"/>
                                    </div>
                                </div>
                                <div class="col-md-1">
                                    <button class="btn btn-primary" v-if="!isShowHeader" @click="showHeaders">Add Header</button>
                                    <button class="btn btn-primary" v-if="isShowHeader" @click="showHeaders">Hide Header</button>
                                </div>
                                <div class="col-md-2">
                                    <div class="row">
                                        <div class="pull-right">
                                            <button class="btn btn-warning" @click="send">Send</button>
                                            <button class="btn btn-primary" @click="save">Save</button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <hr>
                <div class="form-group" v-if="isShowHeader">
                    <label>Headers</label>

                </div>
                <div class="form-group">
                    <label for="body">Body</label>
                    <textarea v-model="endpoint.body" id="body" class="form-control font-terminal" cols="30" rows="10"></textarea>
                </div>
                <div class="form-group">
                    <label for="response">Response</label>
                    <textarea v-model="endpoint.response" id="response" class="form-control font-terminal" cols="30" rows="10" readonly></textarea>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {http} from "../../lib/endpointlib";
    import {HTTPAUTH} from "../../lib/httplib";

    export default {
        name: "Document",
        data() {
            return {
                did: "",
                loading: false,
                endpoint: {},
                isShowHeader: false,
            }
        },
        created() {
            this.did = this.$route.params.did
        },
        methods: {
            showHeaders() {
                this.isShowHeader = !this.isShowHeader
            },
            save() {

            },
            send() {
                var token = "";
                var headers = {
                    'Content-Type': 'application/json',
                    'Authorization': token
                };
                http(headers).post()
            }
        },
        watch: {
            '$route' (to, from) {
                if (to.params) {
                    this.did = to.params.did
                }
            }
        }
    }
</script>

<style scoped>
    .font-terminal {
        font-family: Courier;
    }
</style>