<template>
  <DateInput
    :error="false"
    :error-message="'error message'"
    data-purpose-prop='additional_driver_birthdate'
  />
</template>

<script>
import DateInput from './components/DateInput.vue'

export default {
  name: 'App',
  components: {
    DateInput
  }
}
</script>
