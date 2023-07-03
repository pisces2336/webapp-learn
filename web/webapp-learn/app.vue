<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-4">
        <input v-model="new_kanban.title" placeholder="title" class="form-control">
        <input v-model="new_kanban.body" placeholder="body" class="form-control">
        <button @click="createKanban()" class="btn btn-success">カンバン作成</button>
      </div>
    </div>

    <hr>

    <div class="row mt-3 px-5 h-auto">
      <div v-for="kanban_category in kanbans" class="col-4">
        <h2 v-if="kanbans.indexOf(kanban_category) == 0">未着手: {{ kanban_category.length }}</h2>
        <h2 v-else-if="kanbans.indexOf(kanban_category) == 1">作業中: {{ kanban_category.length }}</h2>
        <h2 v-else-if="kanbans.indexOf(kanban_category) == 2">完了: {{ kanban_category.length }}</h2>

        <TransitionGroup name="list" tag="ul" class="position-relative list-unstyled">
          <li v-for="k in kanban_category" :key="k" class="card w-100 mx-auto mb-3">
            <div class="card-body">
              <h5 class="card-title">{{ k.title }}</h5>
              <p class="card-text">{{ k.body }}</p>
              <div class="d-flex justify-content-between">
                <button v-if="k.category >= 1" @click="moveKanban(k, 'left')" class="btn btn-dark">◀</button>
                <button v-else class="btn btn-secondary" disabled>◀</button>

                <button v-if="k.category <= 1" @click="moveKanban(k, 'right')" class="btn btn-dark">▶</button>
                <button v-else class="btn btn-secondary" disabled>▶</button>
              </div>
            </div>
          </li>
        </TransitionGroup>
      </div>
    </div>
  </div>
</template>
<script setup>
  class Kanban {
    title = '';
    body = '';
    category = 0;
  }

  const new_kanban = ref(new Kanban());
  const kanbans = ref(getCategorizedKanbans());

  const createKanban = () => {
    const k = new_kanban.value
    if (k.title == '' || k.body =='') {
      return
    }
    kanbans.value[0].push(k)
    new_kanban.value = new Kanban()
  }

  const getAllKanbans = async () => {
    const url = "http://api:8080/kanbans";
    const { data } = await useFetch(url);
    return data.value.kanbans;
  }

  const updateKanban = () => {
    kanbans.value = getCategorizedKanbans();
  }

  const deleteKanban = async (k) => {
    const url = "http://api:8080/kanbans/" + String(k.id);
    await useFetch(url, { method: "delete" });
    kanbans.value = getCategorizedKanbans();
  }

  const getCategorizedKanbans = () => {
    const kanbans = getAllKanbans();
    const result = [[], [], []];
    for (let i = 0; i < kanbans.length; i++) {
      let k = kanbans[i];
      result[k.category].push(k);
    }
    return result;
  }

  const moveKanban = (k, direction) => {
    // kanban_list内の対応するオブジェクトを取得
    const idx = kanbans.value[k.category].indexOf(k)
    const kanban = kanbans.value[k.category][idx]

    // 元居たカテゴリから削除
    kanbans.value[kanban.category].splice(idx, 1)

    // カテゴリ移動
    if (direction == 'left') {
      kanban.category--
    } else {
      kanban.category++
    }
    kanbans.value[kanban.category].push(kanban)
  }
</script>