<script lang="ts">
	import { PUBLIC_APP_NAME } from '$env/static/public';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Button } from '$lib/components/ui/button';
	import type { PageData } from './$types';
	import { superForm } from 'sveltekit-superforms';
	import { toast } from 'svelte-sonner';

	export let data: PageData;

	const { form, constraints, enhance, delayed, errors, message } = superForm(data.form);

	$: if ($message?.type === 'success') {
		toast.success($message.text);
	} else if ($message?.type === 'error') {
		toast.error($message.text);
	}
</script>

<svelte:head>
	<title>Register - {PUBLIC_APP_NAME}</title>
</svelte:head>

<div class="flex justify-between gap-4">
	<div class="w-1/2">
		<h2 class="text-3xl font-light">Register</h2>
		<form method="post" use:enhance action="" class="mt-8 space-y-4">
			<div class="flex w-full max-w-sm flex-col gap-1.5">
				<Label for="username">Username</Label>
				<Input
					type="text"
					name="username"
					id="username"
					bind:value={$form.username}
					{...$constraints.username}
				/>
				{#if $errors.username}
					<p class="text-sm text-red-500">{$errors.username}</p>
				{/if}
			</div>
			<div class="flex w-full max-w-sm flex-col gap-1.5">
				<Label for="email">Email</Label>
				<Input
					type="email"
					name="email"
					id="email"
					bind:value={$form.email}
					{...$constraints.email}
				/>
				{#if $errors.email}
					<p class="text-sm text-red-500">{$errors.email}</p>
				{/if}
			</div>
			<div class="flex w-full max-w-sm flex-col gap-1.5">
				<Label for="password">Password</Label>
				<Input
					type="password"
					name="password"
					id="password"
					bind:value={$form.password}
					{...$constraints.password}
				/>
				{#if $errors.password}
					<p class="text-sm text-red-500">{$errors.password}</p>
				{/if}
			</div>
			<Button type="submit" disabled={$delayed}>Register</Button>
		</form>
		<p class="mt-4 text-sm">
			Already have an account? <a href="/auth/login" class="text-blue-500 hover:underline">Login</a>
		</p>
	</div>
	<div class="w-1/2">
		<enhanced:img
			src="../../../lib/assets/illustrations/signup.svg"
			alt="Sign up"
			class="ml-auto h-96 w-auto"
		/>
	</div>
</div>
