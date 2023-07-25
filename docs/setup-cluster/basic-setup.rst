:orphan:

.. _basic-setup:

#########################
 Set Up Determined (WIP)
#########################

.. meta::
   :description: These basic instructions help you get started with Determined by setting up your training environment.

To get started with Determined you'll need to set up your training environment. To set up your
training environment, follow the steps below.

.. note::

   Your training environment can be a local development machine, an on-premise GPU cluster, or cloud
   resources.

.. note::

   Visit the :ref:`Cluster Deployment Guide by Environment <setup-checklists>` for a detailed
   checklist for each environment.

****************************
 Step 1 - Set Up PostgreSQL
****************************

The first step is to set up PostgreSQL on the environment of your choice.

.. note::

   If you are using Kubernetes, you can skip this step. Installing Determined on Kubernetes uses the
   Determined Helm Chart which includes deployment of a PostgreSQL database.

For example, to set up PostgreSQL on Docker, visit Install Determined Using Docker:
:ref:`Preliminary Steps <install-postgres-docker>`.

To validate Step 1, check that PostgreSQL is installed.

****************************************
 Step 2 - Install the Determined Master
****************************************

The next steps is to install the Determined Master.

Install the Determined Master and Agent
=======================================

If the Determined Agent is your compute resource, you'll install the Determined Agent along with the
Determined Master. The preferred method for installing the Agent is to use Linux packages. The
recommended alternative to Linux packages is Docker.

Install Using Linux Packages
----------------------------

To install the Determined Master using Linux packages, visit :ref:`Install Determined Using Linux
Packages--Install the Determined Master and Agent <install-det-linux>`.

.. note::

   This method is required when using Slurm.

Install Using Docker Container (Or Equivalent)
----------------------------------------------

To install the Determined Master using Docker, visit :ref:`Install Determined Using Docker
<install-using-docker>`.

Kubernetes
==========

To install the Determined Master using Kubernetes, visit :ref:`Install Determined on Kubernetes
<install-on-kubernetes>`.

Slurm
=====

To install the Determined Master using Slurm, visit :ref:`Install Determined on Slurm/PBS
<install-on-slurm>`.

********************************
 Step 3 - Set Up TLS (Optional)
********************************

TLS is optional. If you do not require the secure version of HTTP, you can skip this section.

Agent-Based (Where Determined is the Resource Manager)
======================================================

To find out how to set up TLS for Agents, visit :ref:`Transport Security Layer--Agent Configuration
<tls-agent-config>`.

Kubernetes
==========

To set up TLS on Kubernetes, choose one of the following methods:

-  type here
-  type here

Slurm
=====

To set up TLS on Slurm:

*************************************
 Step 4 - Set Up Security (Optional)
*************************************

.. attention::

   SSO is only supported on the Determined Enterprise Edition.

To set up SSO, follow these instructions:

-  x
-  x
-  x

Only changes with Kubernetes.

To validate Step 4, ensure the users can access the Determined cluster.

***********************************
 Step 5 - Set Up Compute Resources
***********************************

Linux Packages
==============

x

Docker
======

x

Slurm
=====

x

Kubernetes
==========

x

*********************************************
 Step 6 - Set Up Monitoring Tools (Optional)
*********************************************

The following monitoring tools are officially supported: Prometheus/Grafana

Prometheus
==========

x

Grafana
=======

x

************
 Next Steps
************

RBAC
====

x

Workspaces
==========

x

Checkpoint Storage
==================

x
