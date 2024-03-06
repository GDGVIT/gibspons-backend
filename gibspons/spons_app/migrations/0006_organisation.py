# Generated by Django 4.2.6 on 2024-01-30 17:40

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('spons_app', '0005_rename_domain_company_industry'),
    ]

    operations = [
        migrations.CreateModel(
            name='Organisation',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=255)),
                ('invite_code', models.CharField(blank=True, max_length=8, unique=True)),
                ('created_at', models.DateTimeField(auto_now_add=True)),
            ],
        ),
    ]
