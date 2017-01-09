<template>

  <div class="course-search">

    <button type="button" @click="toggleSemesterList">{{selectedSemester}}<i class="fa fa-caret-down"></i></button>

    <ul class="semester-list" v-if="showSemesterList">

      <li v-for="semester in $store.state.semesters" @click="selectSemester(semester)">{{semester.name}}</li>

    </ul>

    <input type="text" placeholder="Filter Courses" autocomplete="off" @input="setFilter" />

  </div>

</template>

<script>

  export default {

    name: 'course-search',

    data() {

      return {

        showSemesterList: false

      }

    },

    computed: {

      selectedSemester() {

        let semester = this.$store.state.selectedSemester;

        return semester ? semester.name : 'Select Semester';

      }

    },

    methods: {

      selectSemester(semester) {

        this.$store.dispatch('selectSemester', semester);
        this.showSemesterList = false;

      },

      setFilter(event) {

        this.$store.dispatch('setFilter', event.target.value);

      },

      toggleSemesterList() {

        this.showSemesterList = !this.showSemesterList;

      }

    }

  }

</script>

<style lang="postcss" scoped>

  .course-search {

    display: flex;
    position: relative;

  }

  button {

    background: #eee;
    border: 1px solid #ddd;
    width: 200px;
    margin-right: 10px;
    padding: 10px;
    text-align: left;
    cursor: pointer;
    outline: none;

    i {

      float: right;

    }

  }

  input {

    flex: 1;
    border: 1px solid #ddd;
    font-size: 1rem;
    padding: 10px;
    outline: none;

  }

  .semester-list {

    border: 1px solid #ddd;
    border-top: none;
    top: 100%;
    left: 0;
    margin: 0;
    padding: 0;
    position: absolute;
    width: 198px;

  }

  li {

    padding: 10px;
    cursor: pointer;
    background: #fff;
    list-style: none;

    &:hover {

      background: #f5f5f5;

    }

  }

</style>
