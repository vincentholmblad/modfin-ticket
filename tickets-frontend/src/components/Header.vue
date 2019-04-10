<template>
  <nav
    class="bg-purple-dark text-white pl-6 whitespace-no-wrap text-left shadow relative flex items-center justify-between"
  >
    <h1 class="inline-block text-2xl py-3">ModFin Tickets</h1>
    <div class="relative w-full h-full ml-6">
      <input
        v-model="search"
        type="search"
        class="h-full w-full bg-transparent px-3 pl-12 border-l border-r border-purple text-white block focus:outline-none"
        placeholder="Search tickets..."
      >
      <img
        class="w-4 h-4 mx-6 search-icon opacity-75 pointer-events-none absolute pin-l"
        src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB2ZXJzaW9uPSIxLjEiIGlkPSJDYXBhXzEiIHg9IjBweCIgeT0iMHB4IiB2aWV3Qm94PSIwIDAgNDUxIDQ1MSIgc3R5bGU9ImVuYWJsZS1iYWNrZ3JvdW5kOm5ldyAwIDAgNDUxIDQ1MTsiIHhtbDpzcGFjZT0icHJlc2VydmUiIHdpZHRoPSI1MTIiIGhlaWdodD0iNTEyIiBjbGFzcz0iIj48Zz48Zz4KCTxwYXRoIGQ9Ik00NDcuMDUsNDI4bC0xMDkuNi0xMDkuNmMyOS40LTMzLjgsNDcuMi03Ny45LDQ3LjItMTI2LjFDMzg0LjY1LDg2LjIsMjk4LjM1LDAsMTkyLjM1LDBDODYuMjUsMCwwLjA1LDg2LjMsMC4wNSwxOTIuMyAgIHM4Ni4zLDE5Mi4zLDE5Mi4zLDE5Mi4zYzQ4LjIsMCw5Mi4zLTE3LjgsMTI2LjEtNDcuMkw0MjguMDUsNDQ3YzIuNiwyLjYsNi4xLDQsOS41LDRzNi45LTEuMyw5LjUtNCAgIEM0NTIuMjUsNDQxLjgsNDUyLjI1LDQzMy4yLDQ0Ny4wNSw0Mjh6IE0yNi45NSwxOTIuM2MwLTkxLjIsNzQuMi0xNjUuMywxNjUuMy0xNjUuM2M5MS4yLDAsMTY1LjMsNzQuMiwxNjUuMywxNjUuMyAgIHMtNzQuMSwxNjUuNC0xNjUuMywxNjUuNEMxMDEuMTUsMzU3LjcsMjYuOTUsMjgzLjUsMjYuOTUsMTkyLjN6IiBkYXRhLW9yaWdpbmFsPSIjMDAwMDAwIiBjbGFzcz0iYWN0aXZlLXBhdGgiIHN0eWxlPSJmaWxsOiNGRkZGRkYiIGRhdGEtb2xkX2NvbG9yPSIjZmZmZmZmIj48L3BhdGg+CjwvZz48L2c+IDwvc3ZnPg=="
      >
    </div>
    <button
      class="text-white px-6 h-full font-bold focus:outline-none flex items-center hover:bg-purple"
      @click="createTicket"
    >
      <span class="opacity-75 font-thin text-3xl mr-2">+</span>New ticket
    </button>
    <button
      class="flex-no-shrink border-l border-purple text-white px-6 h-full font-bold text-sm focus:outline-none flex items-center hover:bg-purple"
      @click="signOut"
    >
    {{ $root.userName }}
      <img class="ml-2 w-4 h-4 opacity-75" src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB2ZXJzaW9uPSIxLjEiIGlkPSJDYXBhXzEiIHg9IjBweCIgeT0iMHB4IiB2aWV3Qm94PSIwIDAgNDcxLjIgNDcxLjIiIHN0eWxlPSJlbmFibGUtYmFja2dyb3VuZDpuZXcgMCAwIDQ3MS4yIDQ3MS4yOyIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSIgd2lkdGg9IjUxMiIgaGVpZ2h0PSI1MTIiPjxnPjxnPgoJPGc+CgkJPHBhdGggZD0iTTIyNy42MTksNDQ0LjJoLTEyMi45Yy0zMy40LDAtNjAuNS0yNy4yLTYwLjUtNjAuNVY4Ny41YzAtMzMuNCwyNy4yLTYwLjUsNjAuNS02MC41aDEyNC45YzcuNSwwLDEzLjUtNiwxMy41LTEzLjUgICAgcy02LTEzLjUtMTMuNS0xMy41aC0xMjQuOWMtNDguMywwLTg3LjUsMzkuMy04Ny41LDg3LjV2Mjk2LjJjMCw0OC4zLDM5LjMsODcuNSw4Ny41LDg3LjVoMTIyLjljNy41LDAsMTMuNS02LDEzLjUtMTMuNSAgICBTMjM1LjAxOSw0NDQuMiwyMjcuNjE5LDQ0NC4yeiIgZGF0YS1vcmlnaW5hbD0iIzAwMDAwMCIgY2xhc3M9ImFjdGl2ZS1wYXRoIiBzdHlsZT0iZmlsbDojRkZGRkZGIiBkYXRhLW9sZF9jb2xvcj0iI2ZmZmZmZiI+PC9wYXRoPgoJCTxwYXRoIGQ9Ik00NTAuMDE5LDIyNi4xbC04NS44LTg1LjhjLTUuMy01LjMtMTMuOC01LjMtMTkuMSwwYy01LjMsNS4zLTUuMywxMy44LDAsMTkuMWw2Mi44LDYyLjhoLTI3My45Yy03LjUsMC0xMy41LDYtMTMuNSwxMy41ICAgIHM2LDEzLjUsMTMuNSwxMy41aDI3My45bC02Mi44LDYyLjhjLTUuMyw1LjMtNS4zLDEzLjgsMCwxOS4xYzIuNiwyLjYsNi4xLDQsOS41LDRzNi45LTEuMyw5LjUtNGw4NS44LTg1LjggICAgQzQ1NS4zMTksMjM5LjksNDU1LjMxOSwyMzEuMyw0NTAuMDE5LDIyNi4xeiIgZGF0YS1vcmlnaW5hbD0iIzAwMDAwMCIgY2xhc3M9ImFjdGl2ZS1wYXRoIiBzdHlsZT0iZmlsbDojRkZGRkZGIiBkYXRhLW9sZF9jb2xvcj0iI2ZmZmZmZiI+PC9wYXRoPgoJPC9nPgo8L2c+PC9nPiA8L3N2Zz4=" />
    </button>
  </nav>
</template>

<script>
export default {
  data: () => ({
    search: ""
  }),
  watch: {
    search() {
      this.searchEmit();
    }
  },
  methods: {
    createTicket() {
      this.$bus.emit("open-create-ticket-modal");
    },
    searchEmit() {
      this.$emit("input", this.search);
    },
    signOut() {
      this.$cookie.delete('username');
      this.$root.userName = ''
      this.$router.push("/")
    }
  }
};
</script>

<style>
.search-icon {
  top: 50%;
  margin-top: 0px;
  transform: translateY(-50%);
}
</style>
