<template>

  <div class="detailed-course">

    <span class="course-header">{{course.subject}} {{course.number}} - {{course.title}}</span>

    <ul class="course-sections">

      <li v-for="section in course.sections">

        Section {{section.name}} - {{section.instructor}}

      </li>

    </ul>

  </div>

</template>

<script>

  export default {

    name: 'detailed-course',

    props: ['course']

  }

</script>

<style scoped>

  .detailed-course {



  }

  ul {

    font-size: .75rem;

  }


</style>
