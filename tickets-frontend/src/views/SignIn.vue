<template>
  <div class="about w-full h-full bg-purple-lightest flex items-center justify-center">
    <div class="bg-white rounded-lg shadow-lg w-1/3 overflow-hidden">
      <div class="p-6 border-b border-b border-purple-lightest">
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
