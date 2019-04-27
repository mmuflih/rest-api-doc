<template>
    <div>
        {{fid}}
    </div>
</template>

<script>
    export default {
        name: "Document",
        data() {
            return {
                fid: ""
            }
        },
        created() {
            this.fid = this.$route.params.fid
        },
        watch: {
            '$route' (to, from) {
                if (to.params) {
                    this.fid = to.params.fid
                }
            }
        }
    }
</script>

<style scoped>

</style>