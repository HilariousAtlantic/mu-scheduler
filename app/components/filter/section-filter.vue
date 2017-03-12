<template>

  <div class="section-filter">

    <div v-for="course in $store.getters.selectedCourses" class="course">

      <h3>{{course.subject}} {{course.number}} - {{course.title}}</h3>

      <div v-for="section in course.sections" class="section">

        <input type="checkbox" class="form-checkbox">

        <span>{{section.name}}</span>

      </div>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'section-filter',

    methods: {


    }

  }

</script>

<style scoped>

  .section {
    display: inline-flex;
    align-items: center;
  }

  .section + .section {
    margin-left: 10px;
  }

  input {
    margin-right: 5px;
  }

</style>
