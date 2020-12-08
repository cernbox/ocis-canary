<template>
  <div class="oc-app-bar">
    <section>
      <h3 class="uk-heading-divider oc-mt-s">
        Change Canary Mode
      </h3>
      <div class="oc-mb-s">
        <oc-radio
            v-for="o in versions"
            :key="'option-' + o"
            v-model="version"
            :option="o"
            :label="o"
            class="oc-mr-s"
        />
      </div>
      {{ error }}
    </section>
  </div>
</template>
<script>
export default {
  name: 'App',
  data: function () {
    return {
      versions: ['OCIS', 'Canary', 'Production'],
      version: 'OCIS'
    }
  },
  computed: {
    error () {
      return this.$store.getters['Canary/error']
    }
  },
  watch: {
    version: function (version) {
      this.$store.dispatch('Canary/setCanary', version.toLowerCase())
    }
  }
}
</script>
