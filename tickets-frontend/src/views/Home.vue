<template>
<div class="w-full h-full">
  <div v-if="!$root.userName.length" class="w-full h-full">
    <div class="w-full h-full bg-purple-lightest flex items-center justify-center">
        <div class="w-1/3">
        <div class="box bg-white overflow-hidden w-full">
        <div class="px-12 py-8 border-b border-b border-purple-lightest">
            <img class="mb-6 w-24" src="data:image/svg+xml;utf8;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iaXNvLTg4NTktMSI/Pgo8IS0tIEdlbmVyYXRvcjogQWRvYmUgSWxsdXN0cmF0b3IgMTcuMS4wLCBTVkcgRXhwb3J0IFBsdWctSW4gLiBTVkcgVmVyc2lvbjogNi4wMCBCdWlsZCAwKSAgLS0+CjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgdmVyc2lvbj0iMS4xIiBpZD0iQ2FwYV8xIiB4PSIwcHgiIHk9IjBweCIgdmlld0JveD0iMCAwIDM2MyAzNjMiIHN0eWxlPSJlbmFibGUtYmFja2dyb3VuZDpuZXcgMCAwIDM2MyAzNjM7IiB4bWw6c3BhY2U9InByZXNlcnZlIiB3aWR0aD0iNTEycHgiIGhlaWdodD0iNTEycHgiPgo8cGF0aCBkPSJNMjc3LjczLDk0LjEyM2MwLDEwLjk5Ny04LjAwNiwxNy42ODUtMTMuODUyLDIyLjU5M2MtMi4yMTQsMS44NTktNi4zMzUsNS4yNTEtNi4zMjQsNi41MTggIGMwLjA0LDQuOTctMy45NTYsOC45MzktOC45MjcsOC45MzljLTAuMDI1LDAtMC4wNSwwLTAuMDc1LDBjLTQuOTM2LDAtOC45NTgtMy44NDctOC45OTgtOC43OTIgIGMtMC4wNzktOS43NDcsNy4wMzQtMTUuNTg0LDEyLjc1LTIwLjM4M2M0LjQ4NS0zLjc2Niw3LjQyNi02LjQxNiw3LjQyNi04Ljg0MWMwLTQuOTA5LTMuOTk0LTguOTAzLTguOTAzLTguOTAzICBjLTQuOTExLDAtOC45MDYsMy45OTQtOC45MDYsOC45MDNjMCw0Ljk3MS00LjAyOSw5LTksOXMtOS00LjAyOS05LTljMC0xNC44MzQsMTIuMDY5LTI2LjkwMywyNi45MDQtMjYuOTAzICBDMjY1LjY2MSw2Ny4yNTMsMjc3LjczLDc5LjI4OCwyNzcuNzMsOTQuMTIzeiBNMjQ4LjgwMSwxNDAuNDgxYy00Ljk3MSwwLTguODAxLDQuMDI5LTguODAxLDl2MC4wNjkgIGMwLDQuOTcxLDMuODMxLDguOTY2LDguODAxLDguOTY2czktNC4wNjQsOS05LjAzNVMyNTMuNzcyLDE0MC40ODEsMjQ4LjgwMSwxNDAuNDgxeiBNNjcuMzkyLDIwMy4xNzRjLTQuOTcxLDAtOSw0LjAyOS05LDkgIHM0LjAyOSw5LDksOWgwLjc1YzQuOTcxLDAsOS00LjAyOSw5LTlzLTQuMDI5LTktOS05SDY3LjM5MnogTTk4LjY3MSwyMDMuMTc0Yy00Ljk3MSwwLTksNC4wMjktOSw5czQuMDI5LDksOSw5aDAuNzQ5ICBjNC45NzEsMCw5LTQuMDI5LDktOXMtNC4wMjktOS05LTlIOTguNjcxeiBNMzYzLDU5LjQyNXYxMDEuMzAxYzAsMjMuOTg1LTE5LjIzMiw0My40NDgtNDMuMjE3LDQzLjQ0OEgyMDMuMDY2ICBjLTIuMjgyLDAtNC4xNjEtMC4wMTMtNS43MzMtMC4wNDZjLTEuNjQ3LTAuMDM0LTMuNTAxLTAuMDQ3LTQuMjI0LDAuMDMzYy0wLjc1MywwLjUtMi41OTksMi4xOTEtNC4zNzgsMy44MyAgYy0wLjcwNSwwLjY0OS0xLjUwMywxLjM2My0yLjM2NCwyLjE0OWwtMzMuMDIyLDMwLjA5OGMtMi42MzQsMi40MDMtNi41MzEsMy4wMjUtOS43OTMsMS41ODdjLTMuMjYyLTEuNDM5LTUuNTUyLTQuNjY5LTUuNTUyLTguMjM0ICB2LTk1LjQxN0g0My43MmMtMTQuMDYyLDAtMjUuNzIsMTEuNTIzLTI1LjcyLDI1LjU4M3YxMDEuMzAxYzAsMTQuMDYxLDExLjY1OSwyNS4xMTYsMjUuNzIsMjUuMTE2aDEzMC4zNzQgIGMyLjI0NSwwLDQuMzQ1LDEuMDMxLDYuMDAzLDIuNTQ1TDIwNywzMTcuNTIzdi04NS41MzljMC00Ljk3MSw0LjAyOS05LDktOXM5LDQuMDI5LDksOXYxMDUuOTM4YzAsMy41NjUtMi4wNCw2Ljc0Ny01LjMwMyw4LjE4NiAgYy0xLjE2NywwLjUxNS0yLjMzOSwwLjcxOC0zLjU2NiwwLjcxOGMtMi4yMDQsMC00LjM3OC0wLjkwNS02LjA2OS0yLjQ0OWwtMzkuNDU3LTM2LjIwNEg0My43MmMtMjMuOTg2LDAtNDMuNzItMTkuMTMtNDMuNzItNDMuMTE2ICBWMTYzLjc1N2MwLTIzLjk4NSwxOS43MzQtNDMuNTgzLDQzLjcyLTQzLjU4M0gxMzhWNTkuNDI1YzAtMjMuOTg2LDE5Ljg4NS00My4yNTEsNDMuODcxLTQzLjI1MWgxMzcuOTEzICBDMzQzLjc2OCwxNi4xNzQsMzYzLDM1LjQzOSwzNjMsNTkuNDI1eiBNMzQ1LDU5LjQyNWMwLTE0LjA2MS0xMS4xNTctMjUuMjUxLTI1LjIxNy0yNS4yNTFIMTgxLjg3MSAgQzE2Ny44MSwzNC4xNzQsMTU2LDQ1LjM2NCwxNTYsNTkuNDI1djY5LjgzM3Y4My45MzRsMTguMDk1LTE2LjM1M2MwLjgzOC0wLjc2NSwxLjc3Ny0xLjQ2NSwyLjQ2Mi0yLjA5NyAgYzguMjYzLTcuNjE0LDEwLjM3Ny04LjgzMSwyMS4xNTUtOC42MDljMS40NywwLjAzMSwzLjIyMSwwLjA0Miw1LjM1NCwwLjA0MmgxMTYuNzE3YzE0LjA2LDAsMjUuMjE3LTExLjM4OCwyNS4yMTctMjUuNDQ4VjU5LjQyNXoiIGZpbGw9IiMwMDAwMDAiLz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPGc+CjwvZz4KPC9zdmc+Cg==" />
            <h1 class="text-2xl mb-2">ModFin Tickets</h1>
            <p class="text-grey-dark">System to create and manage tickets</p>
        </div>
        <div class="p-6 pt-0">
            <sign-in-form />
        </div>
        </div>
        <p class="opacity-50 text-sm mt-6">Programmer test by <a href="https://www.vincentholmblad.com" target="_blanc" class="text-black text-underline hover:no-underline">Vincent Holmblad</a></p>
        </div>
    </div>
  </div>
  <div v-else>
    <app-header class="flex-no-grow" id="app-nav" />
    <tickets-list class="flex-grow" id="app-content" />
    <create-ticket-modal />
  </div>
</div>
</template>

<script>
import AppHeader from '@/components/Header.vue'
import CreateTicketModal from '@/components/CreateTicketModal.vue'
import TicketsList from '@/components/TicketsList.vue'
import SignInForm from '@/components/SignInForm.vue'

export default {
  name: 'home',
  components: {
    AppHeader,
    CreateTicketModal,
    TicketsList,
    SignInForm
  }
}
</script>

<style>
#app-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 2;
  height: 60px;
}
#app-content {
  width: 100vw;
  height: 100vh;
  padding-top: 60px;
  overflow: hidden;
}
.box {
  -webkit-box-shadow: 0px 8px 46px rgba(0, 0, 0, .08), 0px 2px 6px rgba(0, 0, 0, .03);
    box-shadow: 0px 8px 46px rgba(0, 0, 0, .08), 0px 2px 6px rgba(0, 0, 0, .03);
    border-radius: 2px;
}
</style>