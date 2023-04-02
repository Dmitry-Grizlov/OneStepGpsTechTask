<template>
    <div>
        <Header></Header>
        <div class="container">
            <section>
                <div class="mx-4 alert alert-danger " role="alert" v-if="!hidden"> {{ error }}
                </div>
            </section>
            <div class="d-flex justify-content-center my-4">
                <div class="input-group-overlay d-lg-flex mx-4">
                    <input class="form-control appended-form-control" id="search-field" type="text"
                        placeholder="Enter the name of the worker" @input="searchDrivers">
                    <div class="input-group-append-overlay"><button class="btn btn-primary">search</button>
                    </div>
                </div>
            </div>

            <div class="mx-4 p-2 bg-white rounded">
                <div v-for="item in items" class="row justify-content-between align-items-center py-2 mx-1 mb-2 rounded"
                    :class="[item.bg]">
                    <div class="col-10">
                        <div class="row align-items-baseline">
                            <p class="h5">Name:
                                <span class="h6 driver-name">{{ item.device.display_name }}</span>
                            </p>
                        </div>
                        <div class="row align-items-baseline">
                            <p class="h5">State:
                                <span class="h6">{{ `${item.device.drive_status} since ${item.device.status_time}`
                                }}</span>
                            </p>
                        </div>
                        <div class="row align-items-baseline">
                            <p class="h5">Address:
                                <span class="h6">{{ item.address }}</span>
                            </p>
                        </div>
                    </div>
                    <div class="d-flex flex-column col-2">
                        <router-link :to="{ name: 'map', params: { id: item.device.device_id } }"
                            class="btn btn-dark mb-2">Show</router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Header from '@/components/Header';
import common from '@/common'

export default {
    components: {
        Header
    },

    data() {
        return {
            items: [],
            hidden: true,
            error: ''
        }
    },

    created() {
        if (!common.isAuthorised()) {
            this.$router.push('/login');
            return;
        }
    },

    async mounted() {
        await this.getData();
    },

    methods: {
        searchDrivers() {
            let field = document.getElementById('search-field');
            let names = document.getElementsByClassName('driver-name');
            let parent;

            if (field.value.length === 0) {
                for (let i = 0; i < names.length; i++) {
                    parent = this.getParent(names[i]);
                    parent.classList.remove("d-none");
                }
            } else {
                for (let i = 0; i < names.length; i++) {
                    parent = this.getParent(names[i]);
                    if (!names[i].innerText.toLowerCase().includes(field.value.toLowerCase())) {
                        parent.classList.add("d-none");
                    }
                    else {
                        parent.classList.remove("d-none");
                    }
                }
            }
        },

        getParent(el) {
            return el.parentElement.parentElement.parentElement.parentElement;
        },

        async getData() {
            let devices = await common.getDvices();
            if (devices == null) {
                this.showError("Could not get data");
                return;
            }

            let address;
            devices.forEach(async x => {
                address = await common.getAddressFrom(x.lat, x.lng)
                x.status_time = new Date(x.status_time * 1000).toLocaleString();
                x.tracker_time = new Date(x.tracker_time * 1000).toLocaleString();

                this.items.push({ device: x, address: address, bg: this.getBg(x.drive_status) })
            });
        },

        showError(data) {
            this.error = data;
            this.hidden = false;
        },

        getBg(item) {
            switch (item) {
                case 'idle':
                    return 'bg-primary'
                case 'off':
                    return 'bg-danger'
                case 'no signal':
                    return 'bg-secondary'
                case 'driving':
                    return 'bg-success'
                default:
                    return 'bg-warning'
            }

        },


    }
}
</script>

<style>
body {
    background: #eee;
}

.input-group-overlay {
    position: relative;
    width: 100%;
}

.form-control {
    display: block;
    width: 100%;
    height: 47px;
    padding: 6px 12px;
    font-size: 14px;
    line-height: 1.42857143;
    color: #555;
    background-color: #fff;
    background-image: none;
    border: 1px solid #ccc;
    border-radius: 4px;
    -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
    box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
}

.input-group-append-overlay {
    right: 0;
}

.input-group-append-overlay,
.input-group-prepend-overlay {
    display: -ms-flexbox;
    display: flex;
    position: absolute;
    top: 0;
    height: 100%;
    z-index: 5;
}

.input-group-text {
    border: 0;
    background-color: transparent;
}

.input-group-text {

    display: flex;
    align-items: center;
    padding: .625rem 1rem;
    margin-bottom: 2px;
    font-size: 1.3rem;
    font-weight: 400;
    line-height: 1.5;
    color: #4b566b;
    text-align: center;
    white-space: nowrap;
    border-radius: .3125rem;
}

.form-control:focus {
    color: #495057;
    background-color: #fff;
    outline: 0;
    box-shadow: 0 0 0 0rem rgba(0, 123, 255, .25) !important;
}


.input-group-text i {

    font-weight: 500;
}
</style>