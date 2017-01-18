<template>

  <div class="filter-list">

    <div class="filter-group">

      <h4>Time Filters</h4>

      <button><i class="fa fa-plus"></i> Filter</button>

    </div>

    <div class="filter">

      <input type="checkbox">

      <span>I want to start after 10:00 AM on Monday, Wednesday, and Friday</span>

    </div>

    <div class="filter">

      <input type="checkbox">

      <span>I want to finish before 4:00 PM on Tuesday and Thursday</span>

    </div>

    <div class="filter-group">

      <h4>Class Filters</h4>

      <button><i class="fa fa-plus"></i> Filter</button>

    </div>

    <div class="filter">

      <input type="checkbox">

      <span>I want less than 3 classes on Monday, Wednesday, and Friday</span>

    </div>

    <div class="filter">

      <input type="checkbox">

      <span>I want more than 1 class on Tuesday and Thursday</span>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'filter-list'

  }

</script>

<style scoped>

  .schedule-selection {

    display: flex;

  }

  .filter-group {

    margin: 15px 0;

    button {

      float: right;
      border: none;
      background: transparent;

    }

  }

  .filter {

    display: flex;
    margin-bottom: 10px;

    input {

      margin-right: 5px;

    }

    span {

      flex: 1;

    }

  }


</style>
