<template>
    <div class="col-md-12">
        <div class="box">
            <div class="overlay" v-if="loading">
                <i class="fa fa-gear fa-spin"></i>
            </div>
            <div class="box-header">
                <div class="box-title"><h4>Title</h4></div>
                <button class="pull-right btn btn-success">Send</button>
            </div>
            <div class="box-body">
                <div class="form-group">
                    <label for="endpoint">Endpoint</label>
                    <input class="form-control" id="endpoint"/>
                </div>
                <div class="form-group">
                    <label for="body">Body</label>
                    <textarea id="body" class="form-control font-terminal" cols="30" rows="10"></textarea>
                </div>
                <div class="form-group">
                    <label for="response">Response</label>
                    <textarea id="response" class="form-control font-terminal" cols="30" rows="10"></textarea>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Document",
        data() {
            return {
                did: "",
                loading: false
            }
        },
        created() {
            this.did = this.$route.params.did
        },
        methods: {

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