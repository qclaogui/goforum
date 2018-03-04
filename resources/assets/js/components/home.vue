<template>
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <!--<div class="card-header">Dashboard</div>-->

                    <div class="card-body">
                        <div class="alert alert-success alert-dismissible fade show">
                            This is a light alert with <a href="#" class="alert-link">an example link</a>

                            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <p class="text-center">{{ body }}</p>
                    </div>
                </div>


                <br>
                <div class="card">
                    <div class="card-body" id="headingOne">
                        <blockquote class="blockquote mb-0" data-toggle="collapse" data-target="#collapseOne">
                            <a href="#">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer posuere erat a ante.</a>
                            <footer class="blockquote-footer">Someone famous in <cite title="Source Title">Source Title</cite></footer>
                        </blockquote>
                    </div>

                    <div id="collapseOne" class="collapse fade" aria-labelledby="headingOne" data-parent="#accordion">
                        <div class="card-body">
                            Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch.
                        </div>
                    </div>
                    <p class="text-right text-muted">Go from qclaogui</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        props: ['body'],

        mounted() {
            console.log('Component mounted.')
        }

    }
</script>