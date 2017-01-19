<template>

  <div class="filter-editor">

    <div class="editor-header">

      <span>Filter Editor</span>

    </div>

    <div class="editor-form">

      <h4>Options</h4>

    </div>

    <div class="editor-result">

      <h4>Result</h4>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'filter-editor'

  }

</script>

<style scoped>

  .filter-editor {

    background: #fff;
    border: 1px solid #ddd;

    h4 {

      margin: 10px 0;

    }

    input {

      padding: 10px;

    }

  }

  .editor-header {

    background: #eee;
    border-bottom: 1px solid #ddd;
    padding: 10px;

  }

  .editor-form {

    padding: 20px;

  }

  .editor-result {

    padding: 20px;
    border-bottom: 1px solid #ddd;

  }

</style>
