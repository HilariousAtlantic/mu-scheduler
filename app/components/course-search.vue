<template>

  <div class="course-search">

    <span @click="toggleSemesterList">{{selectedSemester}}<i class="fa fa-caret-down"></i></span>
    <input type="text" placeholder="Filter Courses" autocomplete="off" @input="setFilter" />

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

        showSemesterList: false

      }

    },

    methods: {

      selectSemester(semester) {

        this.$emit('selectSemester', semester);
        this.showSemesterList = false;

      },

      setFilter(event) {

        this.$emit('changeFilter', event.target.value);

      },

      toggleSemesterList() {

        this.showSemesterList = !this.showSemesterList;

      }

    }

  }

</script>

<style scoped>

  .course-search {

    display: flex;
    position: relative;
    background: #eee;
    border: 1px solid #ddd;
    margin-bottom: 5px;

  }

  span {

    width: 7.5rem;
    padding: 10px;
    cursor: pointer;

    i {

      float: right;

    }

  }

  input {

    flex: 1;
    border: none;
    border-left: 1px solid #ddd;
    font-size: 1rem;
    padding: 10px;
    outline: none;

  }

  .semester-list {

    border: 1px solid #ddd;
    left: -1px;
    top: 100%;
    margin: 0;
    padding: 0;
    position: absolute;

  }

  li {

    padding: 10px;
    border-top: 1px solid #ddd;
    cursor: pointer;
    width: 7.5rem;
    background: #fff;
    list-style: none;

    &:hover {

      background: #f5f5f5;

    }

    &first-of-type {

      border: none;

    }

  }

</style>
