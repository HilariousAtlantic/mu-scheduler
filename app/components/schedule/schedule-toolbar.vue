<template>

  <div class="schedule-toolbar">

    <div class="schedule-options">

      <div class="btn-group">

        <button class="btn"><span class="fa fa-square"></span></button>

        <button class="btn"><span class="fa fa-th-large"></span></button>

        <button class="btn"><span class="fa fa-th"></span></button>

      </div>

      <button class="btn">Schedule Info</button>

    </div>

    <div class="schedule-search">

      <button class="btn" @click="$emit('dec')">

        <span class="fa fa-caret-left"></span>

      </button>

      <input type="text" class="form-control" :value="searchText">

      <button class="btn" @click="$emit('inc')">

        <span class="fa fa-caret-right"></span>

      </button>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'schedule-toolbar',

    props: ['index'],

    computed: {

      searchText() {

        return (this.index+1) + ' of ' + this.$store.getters.filteredSchedules.length;

      }

    }

  }

</script>

<style scoped>

  .schedule-toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .schedule-search button {
    padding: 6px 20px;
  }

  .schedule-search input {
    text-align: center;
  }

</style>
