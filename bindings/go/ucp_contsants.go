/*
 * Copyright (C) 2021, NVIDIA CORPORATION & AFFILIATES. ALL RIGHTS RESERVED.
 * See file LICENSE for terms.
 */

package ucp

// #include <ucp/api/ucp.h>
import "C"

type UcpProtection uint32

const (
	UCP_MEM_MAP_PROT_LOCAL_READ   UcpProtection = C.UCP_MEM_MAP_PROT_LOCAL_READ
	UCP_MEM_MAP_PROT_LOCAL_WRITE  UcpProtection = C.UCP_MEM_MAP_PROT_LOCAL_WRITE
	UCP_MEM_MAP_PROT_REMOTE_READ  UcpProtection = C.UCP_MEM_MAP_PROT_REMOTE_READ
	UCP_MEM_MAP_PROT_REMOTE_WRITE UcpProtection = C.UCP_MEM_MAP_PROT_REMOTE_WRITE
)

type UcpContextAttr uint32

const (
	UCP_ATTR_FIELD_REQUEST_SIZE UcpContextAttr = C.UCP_ATTR_FIELD_REQUEST_SIZE
	UCP_ATTR_FIELD_THREAD_MODE  UcpContextAttr = C.UCP_ATTR_FIELD_THREAD_MODE
	UCP_ATTR_FIELD_MEMORY_TYPES UcpContextAttr = C.UCP_ATTR_FIELD_MEMORY_TYPES
	UCP_ATTR_FIELD_NAME         UcpContextAttr = C.UCP_ATTR_FIELD_NAME
)

type UcpMemAttribute uint32

const (
	UCP_MEM_ATTR_FIELD_ADDRESS  UcpMemAttribute = C.UCP_MEM_ATTR_FIELD_ADDRESS
	UCP_MEM_ATTR_FIELD_LENGTH   UcpMemAttribute = C.UCP_MEM_ATTR_FIELD_LENGTH
	UCP_MEM_ATTR_FIELD_MEM_TYPE UcpMemAttribute = C.UCP_MEM_ATTR_FIELD_MEM_TYPE
)

type UcpWakeupEvent uint32

const (
	UCP_WAKEUP_RMA      UcpWakeupEvent = C.UCP_WAKEUP_RMA
	UCP_WAKEUP_AMO      UcpWakeupEvent = C.UCP_WAKEUP_AMO
	UCP_WAKEUP_TAG_SEND UcpWakeupEvent = C.UCP_WAKEUP_TAG_SEND
	UCP_WAKEUP_TAG_RECV UcpWakeupEvent = C.UCP_WAKEUP_TAG_RECV
	UCP_WAKEUP_TX       UcpWakeupEvent = C.UCP_WAKEUP_TX
	UCP_WAKEUP_RX       UcpWakeupEvent = C.UCP_WAKEUP_RX
	UCP_WAKEUP_EDGE     UcpWakeupEvent = C.UCP_WAKEUP_EDGE
)

type UcpWorkerAttribute uint32

const (
	UCP_WORKER_ATTR_FIELD_THREAD_MODE     UcpWorkerAttribute = C.UCP_WORKER_ATTR_FIELD_THREAD_MODE
	UCP_WORKER_ATTR_FIELD_ADDRESS         UcpWorkerAttribute = C.UCP_WORKER_ATTR_FIELD_ADDRESS
	UCP_WORKER_ATTR_FIELD_ADDRESS_FLAGS   UcpWorkerAttribute = C.UCP_WORKER_ATTR_FIELD_ADDRESS_FLAGS
	UCP_WORKER_ATTR_FIELD_MAX_AM_HEADER   UcpWorkerAttribute = C.UCP_WORKER_ATTR_FIELD_MAX_AM_HEADER
	UCP_WORKER_ATTR_FIELD_NAME            UcpWorkerAttribute = C.UCP_WORKER_ATTR_FIELD_NAME
	UCP_WORKER_ATTR_FIELD_MAX_INFO_STRING UcpWorkerAttribute = C.UCP_WORKER_ATTR_FIELD_MAX_INFO_STRING
)
