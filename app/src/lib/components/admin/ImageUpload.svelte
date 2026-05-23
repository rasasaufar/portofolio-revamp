<script lang="ts">
	import { get } from 'svelte/store';
	import { token } from '$lib/stores/auth';
	import { API_BASE_URL } from '$lib/config';

	let {
		value = '',
		onchange
	}: {
		value?: string;
		onchange?: (url: string) => void;
	} = $props();

	let uploading = $state(false);
	let error = $state('');
	let dragOver = $state(false);

	async function uploadFile(file: File) {
		uploading = true;
		error = '';

		const formData = new FormData();
		formData.append('file', file);

		try {
			const currentToken = get(token);
			const res = await fetch(`${API_BASE_URL}/api/upload`, {
				method: 'POST',
				headers: {
					'Authorization': `Bearer ${currentToken}`
				},
				body: formData
			});

			if (!res.ok) {
				const data = await res.json().catch(() => ({}));
				throw new Error(data.error || 'Upload failed');
			}

			const data = await res.json();
			value = data.url;
			onchange?.(data.url);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Upload failed';
		} finally {
			uploading = false;
		}
	}

	function handleFileInput(e: Event) {
		const input = e.target as HTMLInputElement;
		const file = input.files?.[0];
		if (file) uploadFile(file);
	}

	function handleDrop(e: DragEvent) {
		e.preventDefault();
		dragOver = false;
		const file = e.dataTransfer?.files?.[0];
		if (file) uploadFile(file);
	}

	function handleDragOver(e: DragEvent) {
		e.preventDefault();
		dragOver = true;
	}

	function handleDragLeave() {
		dragOver = false;
	}

	function clearImage() {
		value = '';
		onchange?.('');
	}
</script>

<div class="image-upload">
	{#if value}
		<div class="image-preview">
			<img src={value} alt="Preview" />
			<div class="image-preview-actions">
				<button type="button" class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-danger" onclick={clearImage}>Remove</button>
			</div>
		</div>
		<input type="text" class="***REMOVED***-input" {value} readonly style="margin-top:0.5rem; font-size:0.72rem; font-family:var(--font-mono); color:var(--***REMOVED***-ink-muted);" />
	{:else}
		<div
			class="upload-dropzone"
			class:drag-over={dragOver}
			ondrop={handleDrop}
			ondragover={handleDragOver}
			ondragleave={handleDragLeave}
			role="button"
			tabindex="0"
		>
			{#if uploading}
				<div class="upload-loading">
					<div class="loading-box"></div>
					<p>Uploading...</p>
				</div>
			{:else}
				<p class="drop-label">◇ Drag & drop image here</p>
				<p class="drop-hint">or click to browse</p>
				<input type="file" accept="image/*" onchange={handleFileInput} class="file-input" />
			{/if}
		</div>
	{/if}

	{#if error}
		<p class="upload-error">{error}</p>
	{/if}
</div>

<style>
	.image-upload {
		width: 100%;
	}

	.image-preview {
		position: relative;
		border: 3px solid var(--***REMOVED***-ink);
		box-shadow: 3px 3px 0 0 var(--***REMOVED***-ink);
		overflow: hidden;
		max-width: 300px;
	}

	.image-preview img {
		width: 100%;
		height: 160px;
		object-fit: cover;
		display: block;
	}

	.image-preview-actions {
		position: absolute;
		top: 0.5rem;
		right: 0.5rem;
	}

	.upload-dropzone {
		position: relative;
		border: 3px dashed var(--***REMOVED***-ink);
		padding: 2rem;
		text-align: center;
		cursor: pointer;
		transition: background 0.12s ease, transform 0.12s ease;
		background: var(--***REMOVED***-soft);
	}

	.upload-dropzone:hover,
	.upload-dropzone.drag-over {
		background: var(--***REMOVED***-yellow);
		transform: translate(-2px, -2px);
		box-shadow: 4px 4px 0 0 var(--***REMOVED***-ink);
	}

	.drop-label {
		margin: 0;
		font-family: var(--font-mono);
		font-size: 0.78rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
	}

	.drop-hint {
		margin: 0.3rem 0 0;
		font-family: var(--font-mono);
		font-size: 0.65rem;
		color: var(--***REMOVED***-ink-muted);
		text-transform: uppercase;
	}

	.upload-loading {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.6rem;
	}

	.upload-loading p {
		margin: 0;
		font-family: var(--font-mono);
		font-size: 0.75rem;
		text-transform: uppercase;
	}

	.loading-box {
		width: 12px;
		height: 12px;
		border: 3px solid var(--***REMOVED***-ink);
		animation: spin-box 0.8s linear infinite;
	}

	@keyframes spin-box {
		from { transform: rotate(0deg); }
		to { transform: rotate(360deg); }
	}

	.file-input {
		position: absolute;
		inset: 0;
		opacity: 0;
		cursor: pointer;
	}

	.upload-error {
		font-family: var(--font-mono);
		font-size: 0.72rem;
		color: var(--***REMOVED***-danger);
		margin-top: 0.4rem;
		text-transform: uppercase;
	}
</style>
