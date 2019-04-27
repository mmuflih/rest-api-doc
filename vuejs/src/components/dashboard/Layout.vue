<template>
    <div class="wrapper">
        <header class="main-header">
            <!-- Logo -->
            <router-link :to="{path: '/'}" class="logo">
                <!-- mini logo for sidebar mini 50x50 pixels -->
                <span class="logo-mini"><b>RFD</b></span>
                <!-- logo for regular state and mobile devices -->
                <span class="logo-lg"><b>RestFul Doc</b></span>
            </router-link>
            <!-- Header Navbar: style can be found in header.less -->
            <nav class="navbar navbar-static-top">
                <!-- Sidebar toggle button-->
                <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
                    <span class="sr-only">Toggle navigation</span>
                </a>

                <div class="navbar-custom-menu">
                    <ul class="nav navbar-nav">
                        <!-- Messages: style can be found in dropdown.less-->
                        <!-- Tasks: style can be found in dropdown.less -->
                        <!-- User Account: style can be found in dropdown.less -->
                        <li class="dropdown user user-menu">
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                <img :src="avatar" class="user-image" alt="User Image">
                                <span class="hidden-xs">{{auth.data.name}}</span>
                            </a>
                            <ul class="dropdown-menu">
                                <!-- User image -->
                                <li class="user-header">
                                    <img :src="avatar" class="img-circle" alt="User Image">
                                    <p>
                                        {{auth.data.name}}
                                        <small>Member since {{auth.data.created_at | formatDate}}</small>
                                    </p>
                                </li>
                                <!-- Menu Footer-->
                                <li class="user-footer">
                                    <div class="pull-right">
                                        <a href="#" class="btn btn-default btn-flat" @click="logout">Sign out</a>
                                    </div>
                                </li>
                            </ul>
                        </li>
                        <!-- Control Sidebar Toggle Button -->
                    </ul>
                </div>
            </nav>
        </header>
        <!-- Left side column. contains the logo and sidebar -->
        <aside class="main-sidebar">
            <!-- sidebar: style can be found in sidebar.less -->
            <section class="sidebar">
                <ul class="sidebar-menu">
                    <li class="header">MAIN NAVIGATION</li>
                    <li class="treeview" v-for="(v, k) in folders" :key="k">
                        <a href="#">
                            <i class="fa fa-files-o"></i>
                            <span>{{v.name}}</span>
                            <span class="pull-right-container">
                                <span class="label label-primary pull-right">4</span>
                            </span>
                        </a>
                        <ul class="treeview-menu">
                            <li v-for="(m, ky) in v.child" :key="ky"><a :href="'#/doc/' + m.folder_id"><i class="fa fa-file-o text-green"></i><span> {{m.name}}</span></a></li>
                            <li><a href="#" @click="newInputDoc(v.folder_id)">
                                <i class="fa fa-file-o text-green" v-if="!showNewDoc"></i><span  v-if="!showNewDoc"> Create Document</span>
                                <input type="text" class="form-control" @blur="hideNewDoc" v-if="showNewDoc" v-model="document.name" @keypress="createDocument"/>
                            </a>
                            </li>
                        </ul>
                    </li>
                    <li><a href="#" @click="newInputFolder">
                        <i class="fa fa-file-o text-green" v-if="!showNewFolder"></i><span  v-if="!showNewFolder">Create Folder</span>
                        <input type="text" class="form-control" @blur="hideNewFolder" v-if="showNewFolder" v-model="folder" @keypress="createFolder"/>
                    </a>
                    </li>
                    <li class="header">Settings</li>
                    <li><a href="#"><i class="fa fa-circle-o text-red"></i> <span>Users</span></a></li>
                </ul>
            </section>
            <!-- /.sidebar -->
        </aside>

        <div class="content-wrapper">
            <section class="content-header">
                <h1>
                    Dashboard
                    <small>Control panel</small>
                </h1>
                <ol class="breadcrumb">
                    <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
                    <li class="active">Dashboard</li>
                </ol>
            </section>

            <section class="content">
                <div class="row">
                    <router-view></router-view>
                </div>
            </section>
        </div>
    </div>
</template>

<script>
import {clearToken, getToken} from '@/lib/httplib'
import {HTTP, HTTPAUTH} from "../../lib/httplib";
export default {
    data() {
        return {
            avatar: "./static/img/avatar.jpg",
            auth: {
                data: {
                    name: "",
                    file_id: 0,
                    created_at: "",
                    avatar_url: "",
                }
            },
            showNewFolder: false,
            showNewDoc: false,
            folder: "",
            document: {

            },
            folders: [],
            documents: []
        }
    },
    created: function() {
        this.getCookie()
        this.menus("")
    },
    methods: {
        menus(parentID) {
            HTTPAUTH.get("/api/v1/document?pid=" + parentID).then(res => {
                this.folders = res.data.data
            })
        },
        logout() {
            var co = confirm("Yakin akan logout ?")
            if (co) {
                clearToken()
                location.href = "/"
            }
        },
        getCookie() {
            var tokenData = getToken()
            this.auth = tokenData
        },
        newInputFolder() {
            this.showNewFolder = true
        },
        hideNewFolder() {
            this.showNewFolder = false
        },
        newInputDoc(folderID) {
            this.showNewDoc = true
            this.document.parent_id = folderID
        },
        hideNewDoc() {
            this.showNewDoc = false
        },
        createFolder(e) {
            if (e.keyCode == 13) {
                HTTPAUTH.post("/api/v1/document", {
                    name: this.folder,
                    parent: ""
                }).then(res => {
                    this.$toasted.success("Folder Created").goAway(2000);
                    this.menus("");
                    this.folder = "";
                    this.showNewFolder = false;
                }).catch(eres => {
                    this.$toasted.error("Creating failed").goAway(2000)
                })
            }
        },
        createDocument(e) {
            if (e.keyCode == 13) {
                HTTPAUTH.post("/api/v1/document", this.document).then(res => {
                    this.$toasted.success("Document created").goAway(2000)
                }).catch(eres => {
                    this.$toasted.error("Creating failed").goAway(2000)
                })
            }
        }
    },
    filters: {
        formatDate(date) {
            var d = moment(date)
            return d.format("DD MMM YYYY HH:mm:ss")
        }
    }
}
</script>

<style lang="css">
</style>
