# Generated by Django 4.2.6 on 2024-01-18 18:55

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('spons_app', '0002_remove_event_sponsors_company_created_at_and_more'),
    ]

    operations = [
        migrations.AlterField(
            model_name='event',
            name='brochure',
            field=models.URLField(blank=True, null=True),
        ),
        migrations.AlterField(
            model_name='event',
            name='expected_reg',
            field=models.IntegerField(blank=True, null=True),
        ),
        migrations.AlterField(
            model_name='event',
            name='money_raised',
            field=models.IntegerField(blank=True, null=True),
        ),
        migrations.AlterField(
            model_name='event',
            name='name',
            field=models.CharField(max_length=255, unique=True),
        ),
    ]
