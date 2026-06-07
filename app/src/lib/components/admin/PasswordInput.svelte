<script lang="ts">
	let {
		id,
		value = $bindable(''),
		required = false,
		autocomplete,
		placeholder = '',
		minlength,
		maxlength,
		disabled = false
	}: {
		id: string;
		value: string;
		required?: boolean;
		autocomplete?: 'current-password' | 'new-password';
		placeholder?: string;
		minlength?: number;
		maxlength?: number;
		disabled?: boolean;
	} = $props();

	let visible = $state(false);
</script>

<div class="password-input-wrap">
	<input
		{id}
		class="admin-input password-input"
		type={visible ? 'text' : 'password'}
		bind:value
		{required}
		{autocomplete}
		{placeholder}
		{minlength}
		{maxlength}
		{disabled}
	/>
	<button
		class="password-toggle"
		type="button"
		aria-label={visible ? 'Hide password' : 'Show password'}
		aria-pressed={visible}
		title={visible ? 'Hide password' : 'Show password'}
		onclick={() => { visible = !visible; }}
	>
		{#if visible}
			<svg viewBox="0 0 24 24" aria-hidden="true">
				<path d="M3 3L21 21" />
				<path d="M10.6 10.6A2 2 0 0 0 13.4 13.4" />
				<path d="M8.3 5.6A10.5 10.5 0 0 1 12 5C17 5 20.3 8.1 22 12C21.3 13.6 20.3 15 19 16.1" />
				<path d="M15.5 18.4A11 11 0 0 1 12 19C7 19 3.7 15.9 2 12C2.8 10.1 4.1 8.5 5.6 7.4" />
			</svg>
		{:else}
			<svg viewBox="0 0 24 24" aria-hidden="true">
				<path d="M2 12C3.7 8.1 7 5 12 5S20.3 8.1 22 12C20.3 15.9 17 19 12 19S3.7 15.9 2 12Z" />
				<path d="M12 15A3 3 0 1 0 12 9A3 3 0 0 0 12 15Z" />
			</svg>
		{/if}
	</button>
</div>

<style>
	.password-input-wrap {
		position: relative;
		width: 100%;
	}

	.password-input {
		width: 100%;
		box-sizing: border-box;
		padding-right: 3.25rem;
	}

	.password-toggle {
		position: absolute;
		top: 50%;
		right: 0.45rem;
		display: inline-flex;
		align-items: center;
		justify-content: center;
		width: 2.1rem;
		height: 2.1rem;
		padding: 0;
		color: var(--admin-ink);
		background: var(--admin-white);
		border: 2px solid var(--admin-ink);
		box-shadow: 2px 2px 0 0 var(--admin-ink);
		cursor: pointer;
		transform: translateY(-50%);
		transition: transform 0.1s ease, box-shadow 0.1s ease, background 0.1s ease;
	}

	.password-toggle:hover {
		background: var(--admin-yellow);
		transform: translate(-1px, calc(-50% - 1px));
		box-shadow: 3px 3px 0 0 var(--admin-ink);
	}

	.password-toggle:active {
		transform: translate(1px, calc(-50% + 1px));
		box-shadow: 1px 1px 0 0 var(--admin-ink);
	}

	.password-toggle:focus-visible {
		outline: 3px solid var(--admin-ink);
		outline-offset: 2px;
	}

	.password-toggle svg {
		width: 1.1rem;
		height: 1.1rem;
		fill: none;
		stroke: currentColor;
		stroke-width: 2.2;
		stroke-linecap: round;
		stroke-linejoin: round;
	}
</style>
