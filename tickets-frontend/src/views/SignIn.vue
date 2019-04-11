<template>
  <div class="w-full h-full bg-purple-lightest flex items-center justify-center">
    <div class="box bg-white w-1/3 overflow-hidden">
      <div class="px-12 py-6 border-b border-b border-purple-lightest">
        <h1 class="text-2xl mb-2">ModFin Tickets</h1>
        <p class="text-grey-dark">System to create and manage tickets</p>
      </div>
        <sign-in-form />
    </div>
  </div>
</template>
<script>
import SignInForm from '@/components/SignInForm.vue'
export default {
  components: {
    SignInForm
  }
}
</script>
<style>
.box {
  -webkit-box-shadow: 0px 8px 46px rgba(0, 0, 0, .08), 0px 2px 6px rgba(0, 0, 0, .03);
    box-shadow: 0px 8px 46px rgba(0, 0, 0, .08), 0px 2px 6px rgba(0, 0, 0, .03);
    border-radius: 2px;
}
</style>