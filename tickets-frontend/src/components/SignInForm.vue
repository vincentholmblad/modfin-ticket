<template>
  <form class="w-full" @submit.prevent="signIn()">
    <input
      type="text"
      v-model="name"
      placeholder="Enter your name here"
      ref="name"
      class="bg-grey-white text-xl appearance-none border-2 border-none w-full py-6 px-6 text-grey-darker leading-tight focus:outline-none focus:bg-white focus:border-purple text-center"
    >
    <div class="p-2">
      <button
        @click="signIn"
        class="shadow bg-purple hover:bg-purple-light rounded focus:outline-none focus:outline-none text-white font-bold py-6 px-6 w-full"
        type="button"
      >Continue</button>
    </div>
  </form>
</template>

<script>
export default {
  data: () => ({
    name: ""
  }),
  mounted() {
    this.$refs.name.focus();
    if (this.$cookie.get("username")) {
      this.name = this.$cookie.get("username");
      this.signIn();
    }
  },
  methods: {
    signIn() {
      if (this.name.length) {
        this.$cookie.set("username", this.name);
        this.$root.userName = this.name;
        this.$router.push("/tickets");
      }
    }
  }
};
</script>

<style>
</style>
