<template>

  <div class="course">

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

    name: 'course',

    props: ['course']

  }

</script>

<style scoped>

  .course {



  }

  ul {

    font-size: .75rem;

  }


</style>
