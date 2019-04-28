<template>
    <div class="wrapper">
        <div class="ov">
        </div>
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
                        <li class="dropdown notifications-menu">
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                <span>{{project.name}}</span>
                                <span class="label label-warning">{{projects.length}}</span>
                            </a>
                            <ul class="dropdown-menu">
                                <li class="header">You have {{projects.length}} projects</li>
                                <li>
                                    <!-- inner menu: contains the actual data -->
                                    <ul class="menu">
                                        <li v-for="(pv, pk) in projects" :key="pk">
                                            <a href="#" @click="selectProject(pv)">
                                                <i class="fa fa-gear"></i> {{pv.name}}
                                            </a>
                                        </li>
                                    </ul>
                                </li>
                                <li class="footer">
                                    <a href="#">Add Project</a>
                                </li>
                            </ul>
                        </li>
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
                                <span class="label label-primary pull-right">{{v.child.length}}</span>
                            </span>
                        </a>
                        <ul class="treeview-menu">
                            <li v-for="(m, ky) in v.child" :key="ky">
                                <a :href="'#/doc/' + v.folder_id + '/' + m.folder_id">
                                    <span class="label label-primary" v-if="!m.method">GET</span>
                                    <span class="label label-primary" v-if="m.method">{{m.method}}</span>
                                    <span> {{m.name}}</span>
                                </a>
                            </li>
                            <li>
                                <a href="#" @click="newInputDoc(v.folder_id)">
                                    <span  v-if="!showNewDoc"> Create Document</span>
                                    <input type="text" class="form-control" @blur="hideNewDoc" v-if="showNewDoc" v-model="document.name" @keypress="createDocument"/>
                                </a>
                            </li>
                        </ul>
                    </li>
                    <li><a href="#" @click="newInputFolder">
                        <span  v-if="!showNewFolder">Create Folder</span>
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
    import {HTTPAUTH} from "../../lib/httplib";

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
            documents: [],
            projects: [
                {id: "prj001", name: "Project 01"},
                {id: "prj002", name: "Project 02"},
                {id: "prj003", name: "Project 03"},
            ],
            project: null,
            loading: true
        }
    },
    created: function() {
        this.getCookie();
        this.menus("");
        this.getProjectCookie();
    },
    methods: {
        getProjectCookie() {
            let project = this.$cookie.get('project');
            if (!project) {
                this.project = {id: '000', name: 'Select Project'};
                return
            }
            this.project = JSON.parse(project);
        },
        selectProject(project) {
            this.$cookie.set('project', JSON.stringify(project));
            this.project = project;
        },
        menus(parentID) {
            HTTPAUTH.get("/api/v1/document?pid=" + parentID).then(res => {
                this.folders = res.data.data
            })
        },
        logout() {
            var co = confirm("Yakin akan logout ?")
            if (co) {
                this.redirectLogin()
            }
        },
        redirectLogin() {
            clearToken();
            location.href = "/"
        },
        getCookie() {
            this.auth = getToken()
        },
        newInputFolder() {
            this.showNewFolder = true
        },
        hideNewFolder() {
            this.showNewFolder = false
        },
        newInputDoc(folderID) {
            this.showNewDoc = true;
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
                    this.menus("");
                    this.document = {};
                    this.showNewDoc = false;
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
    .ov {
        position: fixed; /* Sit on top of the page content */
        display: none; /* Hidden by default */
        width: 100%; /* Full width (cover the whole page) */
        height: 100%; /* Full height (cover the whole page) */
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(0,0,0,0.5); /* Black background with opacity */
        z-index: 2; /* Specify a stack order in case you're using a different order for other elements */
        cursor: pointer; /* Add a pointer on hover */
    }
</style>
