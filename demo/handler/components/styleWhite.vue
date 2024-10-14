<template>
    <el-container>
        <el-header v-if="!$route.meta.default_menu">
            <img class="logo" :src="logo" alt="" style="height: 40px" />
            <div style="display: flex; flex-flow: row nowrap; align-items: center">
                <div style="
            margin-right: 40px;
            font-size: 14px;
            display: flex;
            flex-flow: row nowrap;
            align-items: center;
          ">
                    <el-icon style="margin-right: 5px; font-size: 16px">
                        <Message />
                    </el-icon>
                    <div>通知</div>
                </div>
                <el-avatar :src="user.avatar" style="margin-right: 10px; width: 29px; height: 29px">
                    <el-icon style="font-size: 20px">
                        <UserFilled />
                    </el-icon>
                </el-avatar>
                <div style="margin-right: 10px; font-size: 14px">{{ user.nick_name }}</div>
                <el-button link @click="logout" style="color: #9b9faa; font-size: 12px">退出</el-button>
            </div>
        </el-header>
        <el-container class="main">
            <el-aside v-if="!$route.meta.default_menu" @mouseenter="btnCollapseShow = true"
                @mouseleave="() => { btnCollapseShow = false; isHideActive = false }"
                :class="[isHide ? 'hideAside' : '', isHideActive ? 'hideAsideActive' : '']">
                <el-menu :default-active="$route.name" @open="handleOpen" @close="handleClose" :collapse="isCollapse"
                    unique-opened router>
                    <MenuTree :menus="menuList" />
                    <img class="btnCollapseShow" @click="isCollapse = !isCollapse" :src="isCollapse ? toRight : toLeft"
                        alt="" v-if="btnCollapseShow && !isHide" />
                </el-menu>
                <img class="btnCollapseShow" style="position: fixed; left: 0px; z-index: 4; transform: rotateY(180deg)"
                    @click="isHideActive = !isHideActive" :src="toLeft" alt="" v-if="isHide && !isHideActive" />
            </el-aside>
            <el-main>
                <router-view v-slot="{ Component }">
                    <keep-alive>
                        <component :is="Component" :key="$route.path" />
                    </keep-alive>
                </router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
import { ref, onBeforeMount, getCurrentInstance, h } from 'vue'

export default {
    setup() {
        const { proxy } = getCurrentInstance()
        const userStore = proxy.$useUserStore()
        const routerStore = proxy.$useRouterStore()
        
        const menuList = ref([])
        routerStore.loadServerRouter().then(res => {
            menuList.value = res
        })

        const user = ref({})

        const logout = () => {
            userStore.logOut()
        }

        const toRight =
            'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAABQCAYAAAAQq/UNAAAAAXNSR0IArs4c6QAAAotJREFUWEft1j1oFEEUB/D/3EUTo4VEECSCRgKxEMR4FooRGwu1EkHQ2Nv4gZVEtDEono2oYCU2qYKdhY2NH40EFGw1YKHYRzA7H7fvb+YgYe+yl92dIWCR65ab95uZt2/fjELkT0XGYwPARg4QmQOSfVGFRHIoCrDkeBSgdXojCkhM600wYMgDdPwaBPjsG8e3JE9WBkjWtJHnULjie0klYGGBQ/0DfEHw3HIjKgWQVNZiUsBHAHdlu1ghoB1PQdgkeCiv/fUEtOYYFR8DPL1W31wF+OWaFm5S5D6AgaKm2wGQ3KStvARwuSgwN4mJTWfA8sEdr1FbXiXlWdmZO1bwh9zRZ2VeAduDAG3ShwRuVQ1ub4FkXTv+BrkzCHCOx1siH0OC2yswjlMi8iAYSEw6C+BCMLBo0s8KGA8GEiO/AA5HAOlfAINZwBiL/v7NpUyVmLQFoL482jmH2Vev0Th8EPvHRgsRD7B71Pz8D3ya+1IKyQU8WBbpCZRF1gSyyLGjDewb2bMqJ+sLlMnD+iSxzMwrHam7DioXkrbi/GEZU8qrvoXC+s0MUNrIT4K7qwRlx6rEyhzII+GASWeqnETdEylteY2Up8ErIDmqrXwPBnxgYuU9yBMhSPt0ThI3oer1d/7+UxVZOd6XWtudpdY2HQz4QJ9QUJoEtpSF8m4oe7WR26qGiyS2FUE970gktyY2PaugzgM4A+Rjhbc0vwKPaZdeAmtTAEc6Srloidn/SQ4YJ00S11f6QRVgeaxxvCcid/1zqS10T+KvgtrRF99EENAuPscJiHwIBtqISb9FAdqmT6KAxHAyCrCWjShgcZHDUQDJwVigHgUEV2Lwx5T33WxsIbAf/F9v4R9oM4yenR32gAAAAABJRU5ErkJggg=='
        const toLeft =
            'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAABQCAYAAAAQq/UNAAAAAXNSR0IArs4c6QAAAotJREFUWEft1r9rFEEUB/Dv5HLmjBYSQZBYKARsbDxULIzYWKggiKho7G38gZVEtFEUz0ZUsBKbVP4DNjb+AAVJinT+CIoo9hHMzo/b9zWzcEcu2c3uzWAjufb2febN27dvRiHypyLjsQZgrQaIrAHJwahGIjkSBViyGQVonV6JAhLTfhEMGHIXHWeDAF994/iS5KG+AZID2sgTKFzws6QvYH6eI0MNPiV4ojOIKgEklbWYEPA+wK1Lp1gpoB0PQ9giuDtv/BUCWnMnFR8APLLa3FwB+HRNG1cpcgdAo2zo9gAk69rKMwDnywJzi5jYdAqsHtzzGrXlRVIeV125J4Pf5OZBK3MK2BQEaJPeI3Ct3+BsCyRr2vEXyC1BgHM80BZ5GxKcZWAcJ0XkbjCQmPQ5gNPBwIJJZxTQDAYSIz8BjkYA6R8Aw2WAMRZDQ+tWPKYSk7YB1FYDPn6aw/TMLM6cOo56vd7zqAdYJXj/vibGxnbkZlAIdFYuCs76oCiDKsGFwNdv3/Hu/TRWW7n7NeZlEA14PWoLnfSqIP/2NS7PJLeRtBXnD8uYVq70LRQtoLSRHwS3lWVQCCRWPoDcGw6YdKqfk2j5QkpbXiLlUXAGJMe0lS/BgA9MrLwGeTAEyU7nJHHjqlZ75e8//SLd431xtN1YHG23gwEf6AsKSovA+qpQ3g1luzZyXQ3gLImNZVDhHYnkhsSmxxTUSQBHgXys9JbmM/CYduk5cGASYM9orgR0tkGyYZy0SFzuzsSyPeb9bxxvicjNbCqHAP4qqB19840HAVnzOY5D5E0wkCEm/RwFaJs+jAISw4kowFruiQIWFjgaBZAcjgVqUUBwJy7t3rUMAj/n/6yIfwEAJIyenAjZLwAAAABJRU5ErkJggg=='


        const isCollapse = ref(false)
        const isHide = ref(false)
        const isHideActive = ref(false)
        const btnCollapseShow = ref(false)

        const logo = ref(window.adminBox.Logo ? window.adminBox.Logo : './images/logo.png')

        onBeforeMount(async () => {
            user.value = await userStore.userInfo()
            const menuList = await proxy.$useRouterStore().loadServerRouter()
            setCollase()
        })

        const handleOpen = (key, keyPath) => { }
        const handleClose = (key, keyPath) => { }

        const setCollase = () => {
            if (document.body.clientWidth < 1200 && document.body.clientWidth > 1000) {
                isCollapse.value = true
            } else {
                isCollapse.value = false
            }

            if (document.body.clientWidth <= 1000) {
                isHide.value = true
            } else {
                isHide.value = false
            }
        }

        window.onresize = () => {
            setCollase()
        }

        return {
            user,
            logout,
            toRight,
            toLeft,
            isCollapse,
            isHide,
            isHideActive,
            btnCollapseShow,
            logo,
            handleOpen,
            handleClose,
            menuList
        }
    }
}
</script>

<style>
/* 保持原有样式不变 */
html,
body {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    background-color: #eff2f8;
    width: 100%;
    height: 100%;
}

#app {
    min-width: 750px;
    overflow-x: scroll;
}

.el-header,
.el-aside,
.el-main {
    margin: 5px;
    padding: 0px;
    overflow: hidden;
    overflow-y: scroll;
    border-radius: 12px;
}

.el-header {
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-between;
}

.el-aside,
.el-main {
    width: auto;
    height: calc(100vh - 50px - 10px);
}

.el-menu {
    min-width: 220px;
    max-width: 240px;
    min-height: 100%;
    border-radius: 12px;
    padding: 10px 10px;
    border-right: 0;
}

.el-menu--collapse {
    width: unset;
    min-width: unset;

    .el-menu-item {
        box-sizing: content-box;
        padding: 12px;
        line-height: 50px;
    }

    .el-menu-item:hover {
        background: #eff2f8;
    }

    .el-menu-item.is-active {
        color: #2473ff;
        background: #eff2f8;
    }
}

.el-menu.el-menu--inline {
    padding: auto 0px;
}

.el-sub-menu__title .el-sub-menu .el-menu-item,
.el-menu-item {
    border-radius: 6px;
    height: auto;
    line-height: 50px;
}

.el-menu-item:hover {
    background: linear-gradient(to right, #eff2f8, #fff);
}

.el-menu-item.is-active {
    color: #fff;
    background: linear-gradient(to right, #2473ff, #fff);
}

.btnCollapseShow {
    position: absolute;
    top: calc(50% - 80px);
    right: 0;
    cursor: pointer;
}

.hideAside {
    position: fixed;
    z-index: 50;
    left: -230px;
    transition: left 0.35s;
    box-shadow: 0 0 15px 5px rgba(0, 0, 0, 0.2);
}

.hideAsideActive {
    left: 0px;
}

.app-content {
    height: 100%;
}

.main-content,
.content-box {
    box-sizing: border-box;
    background-color: #fff;
    padding: 20px;
    border-radius: 12px;
    height: 100%;
    width: 100%;
    overflow-y: scroll;
    position: relative;
}

:root {

    .el-button {
        border-radius: 3px;
    }
}
</style>
