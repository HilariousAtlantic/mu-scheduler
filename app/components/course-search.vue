<template>

  <div class="course-search">

    <span @click="toggleSemesterList">{{selectedSemester}}<i class="fa fa-caret-down"></i></span>
    <input type="text" placeholder="Filter Courses" autocomplete="off" @input="$emit('filter', $event.target.value)" />

    <ul class="semester-list" v-if="showSemesterList">
      <li v-for="semester in semesters" @click="selectSemester(semester)">{{semester}}</li>
    </ul>

  </div>

</template>

<script>

  export default {

    name: 'course-search',

    props: ['semesters', 'selectedSemester'],

    data() {

      return {

        searchTerm: '',
        showSemesterList: false

      }

    },

    methods: {

      selectSemester(semester) {

        this.$emit('semester', semester);
        this.showSemesterList = false;

      },

      toggleSemesterList() {

        this.showSemesterList = !this.showSemesterList;

      }

    }

  }

</script>

<style>

</style>
