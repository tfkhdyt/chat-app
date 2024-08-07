<script lang="ts">
	import { PUBLIC_APP_NAME } from '$env/static/public';
</script>

<main class="min-h-dvh bg-gray-900">
	<header class="w-full bg-sky-600">
		<div class="container mx-auto max-w-5xl px-0 pb-44 pt-8">
			<h1 class="text-sm font-medium uppercase text-white">Welcome to {PUBLIC_APP_NAME}</h1>
		</div>
	</header>
	<div class="relative">
		<div
			class="container absolute inset-x-0 -top-24 mx-auto max-w-5xl rounded-sm bg-white px-14 py-16 text-gray-600 shadow-2xl"
		>
			<slot />
		</div>
	</div>
</main>
